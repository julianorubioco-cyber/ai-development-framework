from pathlib import Path
import re
import sys

root = Path(__file__).resolve().parents[1]
errors = []
skills_root = root / "skills"
required = {"name", "description", "argument-hint"}

for skill_dir in sorted(p for p in skills_root.iterdir() if p.is_dir()):
    path = skill_dir / "SKILL.md"
    if not path.exists():
        errors.append(f"Ausente: {path}")
        continue

    text = path.read_text(encoding="utf-8")
    match = re.match(r"^---\n(.*?)\n---\n", text, flags=re.S)
    if not match:
        errors.append(f"Frontmatter inválido: {path}")
        continue

    fields = {}
    for line in match.group(1).splitlines():
        if ":" in line:
            key, value = line.split(":", 1)
            fields[key.strip()] = value.strip()

    missing = required - fields.keys()
    if missing:
        errors.append(f"Campos ausentes em {path}: {sorted(missing)}")

    if fields.get("name") != skill_dir.name:
        errors.append(f"Nome divergente em {path}: {fields.get('name')}")

    if "<raiz-do-projeto>/.claude/" not in text and skill_dir.name in {"memory"}:
        errors.append(f"Contrato de isolamento ausente em {path}")

for script in ["install.ps1", "uninstall.ps1", "install.sh", "uninstall.sh"]:
    if not (root / "scripts" / script).exists():
        errors.append(f"Script ausente: {script}")

if errors:
    print("\n".join(errors))
    sys.exit(1)

print("Validação concluída com sucesso.")
