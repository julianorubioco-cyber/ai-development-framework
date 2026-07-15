from __future__ import annotations

import json
import os
import platform
import shutil
import subprocess
import sys
from datetime import datetime, timezone
from importlib import resources
from pathlib import Path
from typing import Any

FRAMEWORK_NAME = "AI Development Framework"


def detect_os() -> str:
    name = platform.system().lower()
    if name == "windows":
        return "windows"
    if name == "darwin":
        return "macos"
    if name == "linux":
        return "linux"
    return name or "unknown"


def claude_root() -> Path:
    return Path.home() / ".claude"


def skills_target() -> Path:
    return claude_root() / "skills"


def manifest_path() -> Path:
    return claude_root() / "adf-install-manifest.json"


def backups_root() -> Path:
    return claude_root() / "adf-backups"


def asset_root() -> Path:
    package_root = resources.files("adf_cli")
    return Path(str(package_root.joinpath("assets")))


def skill_source() -> Path:
    return asset_root() / "skills"


def template_source() -> Path:
    return asset_root() / "templates" / "project-workspace" / ".claude"


def installed_skill_names() -> list[str]:
    source = skill_source()
    return sorted(p.name for p in source.iterdir() if p.is_dir())


def load_manifest() -> dict[str, Any] | None:
    path = manifest_path()
    if not path.exists():
        return None
    return json.loads(path.read_text(encoding="utf-8-sig"))


def save_manifest(data: dict[str, Any]) -> None:
    path = manifest_path()
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(data, indent=2, ensure_ascii=False), encoding="utf-8")


def install(version: str, dry_run: bool = False) -> dict[str, Any]:
    target = skills_target()
    source = skill_source()
    names = installed_skill_names()
    timestamp = datetime.now().strftime("%Y%m%d-%H%M%S")
    backup = backups_root() / timestamp
    existing = [name for name in names if (target / name).exists()]

    result = {
        "os": detect_os(),
        "target": str(target),
        "skills": names,
        "backup": str(backup) if existing else None,
        "dry_run": dry_run,
    }
    if dry_run:
        return result

    target.mkdir(parents=True, exist_ok=True)
    if existing:
        backup.mkdir(parents=True, exist_ok=True)
        for name in existing:
            shutil.copytree(target / name, backup / name, dirs_exist_ok=True)

    for name in names:
        destination = target / name
        if destination.exists():
            shutil.rmtree(destination)
        shutil.copytree(source / name, destination)

    save_manifest({
        "framework": FRAMEWORK_NAME,
        "version": version,
        "installed_at": datetime.now(timezone.utc).isoformat(),
        "os": detect_os(),
        "skills": names,
        "backup_path": str(backup) if existing else None,
    })
    return result


def uninstall(restore_backup: bool = False, dry_run: bool = False) -> dict[str, Any]:
    manifest = load_manifest()
    if not manifest:
        raise RuntimeError(
            f"Manifesto não encontrado em {manifest_path()}. "
            "Execute 'adf doctor' para diagnóstico."
        )

    target = skills_target()
    names = list(manifest.get("skills", []))
    backup_value = manifest.get("backup_path")
    backup = Path(backup_value) if backup_value else None

    result = {
        "skills": names,
        "restore_backup": restore_backup,
        "backup": str(backup) if backup else None,
        "dry_run": dry_run,
    }
    if dry_run:
        return result

    for name in names:
        path = target / name
        if path.exists():
            shutil.rmtree(path)

    if restore_backup and backup and backup.exists():
        for item in backup.iterdir():
            if item.is_dir():
                destination = target / item.name
                if destination.exists():
                    shutil.rmtree(destination)
                shutil.copytree(item, destination)

    manifest_path().unlink(missing_ok=True)
    return result


def find_project_root(start: Path | None = None) -> Path:
    current = (start or Path.cwd()).resolve()

    for candidate in [current, *current.parents]:
        if (candidate / ".git").exists():
            return candidate

    project_markers = {
        "package.json", "pyproject.toml", "Cargo.toml", "go.mod",
        "pom.xml", "build.gradle", "composer.json", "pubspec.yaml",
        ".project", "README.md",
    }
    if any((current / marker).exists() for marker in project_markers):
        return current

    # Empty folder is valid when user explicitly runs `adf init`.
    return current


def init_workspace(project: Path | None = None, dry_run: bool = False) -> dict[str, Any]:
    root = find_project_root(project)
    source = template_source()
    target = root / ".claude"

    files_created: list[str] = []
    files_preserved: list[str] = []

    for source_file in source.rglob("*"):
        if not source_file.is_file():
            continue
        relative = source_file.relative_to(source)
        destination = target / relative
        if destination.exists():
            files_preserved.append(str(relative))
            continue
        files_created.append(str(relative))
        if not dry_run:
            destination.parent.mkdir(parents=True, exist_ok=True)
            shutil.copy2(source_file, destination)

    return {
        "project_root": str(root),
        "workspace": str(target),
        "created": files_created,
        "preserved": files_preserved,
        "dry_run": dry_run,
    }


def command_exists(command: str) -> bool:
    return shutil.which(command) is not None


def doctor() -> dict[str, Any]:
    manifest = load_manifest()
    expected = installed_skill_names()
    target = skills_target()
    missing = [name for name in expected if not (target / name / "SKILL.md").exists()]
    malformed: list[str] = []

    for name in expected:
        path = target / name / "SKILL.md"
        if not path.exists():
            continue
        text = path.read_text(encoding="utf-8")
        if not text.startswith("---\n") or "name:" not in text or "description:" not in text:
            malformed.append(name)

    return {
        "os": detect_os(),
        "python": sys.version.split()[0],
        "git": command_exists("git"),
        "claude": command_exists("claude"),
        "manifest": bool(manifest),
        "installed_version": manifest.get("version") if manifest else None,
        "skills_target": str(target),
        "missing_skills": missing,
        "malformed_skills": malformed,
        "healthy": bool(manifest) and not missing and not malformed,
    }


def git_update(repository: Path | None = None) -> dict[str, Any]:
    root = find_project_root(repository)
    if not (root / ".git").exists():
        raise RuntimeError("A pasta atual não é um clone Git do framework.")

    status = subprocess.run(
        ["git", "-C", str(root), "status", "--porcelain"],
        check=True, capture_output=True, text=True
    ).stdout.strip()
    if status:
        raise RuntimeError(
            "Há alterações locais não salvas. Faça commit ou descarte antes de atualizar."
        )

    subprocess.run(["git", "-C", str(root), "pull", "--ff-only"], check=True)
    return {"repository": str(root), "updated": True}
