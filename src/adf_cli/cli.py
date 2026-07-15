from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

from . import __version__
from .core import doctor, git_update, init_workspace, install, uninstall


def print_result(data: dict) -> None:
    print(json.dumps(data, indent=2, ensure_ascii=False))


def parser() -> argparse.ArgumentParser:
    root = argparse.ArgumentParser(
        prog="adf",
        description="AI Development Framework CLI"
    )
    root.add_argument("--version", action="version", version=f"ADF {__version__}")
    sub = root.add_subparsers(dest="command", required=True)

    install_cmd = sub.add_parser("install", help="Instala as Skills globalmente")
    install_cmd.add_argument("--dry-run", action="store_true")

    init_cmd = sub.add_parser("init", help="Inicializa .claude/ no projeto atual")
    init_cmd.add_argument("path", nargs="?", type=Path)
    init_cmd.add_argument("--dry-run", action="store_true")

    doctor_cmd = sub.add_parser("doctor", help="Diagnostica a instalação")

    update_cmd = sub.add_parser("update", help="Atualiza um clone Git e reinstala")
    update_cmd.add_argument("path", nargs="?", type=Path)

    uninstall_cmd = sub.add_parser("uninstall", help="Remove as Skills do ADF")
    uninstall_cmd.add_argument("--restore-backup", action="store_true")
    uninstall_cmd.add_argument("--dry-run", action="store_true")

    return root


def main() -> None:
    args = parser().parse_args()
    try:
        if args.command == "install":
            result = install(__version__, dry_run=args.dry_run)
        elif args.command == "init":
            result = init_workspace(args.path, dry_run=args.dry_run)
        elif args.command == "doctor":
            result = doctor()
        elif args.command == "update":
            git_update(args.path)
            result = install(__version__)
        elif args.command == "uninstall":
            result = uninstall(
                restore_backup=args.restore_backup,
                dry_run=args.dry_run,
            )
        else:
            raise RuntimeError("Comando desconhecido.")
        print_result(result)
    except Exception as exc:
        print(f"Erro: {exc}", file=sys.stderr)
        raise SystemExit(1)


if __name__ == "__main__":
    main()
