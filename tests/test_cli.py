from pathlib import Path
import os
import tempfile
import unittest
from unittest.mock import patch

from adf_cli.core import detect_os, find_project_root, init_workspace


class CliTests(unittest.TestCase):
    def test_detect_os_returns_value(self):
        self.assertTrue(detect_os())

    def test_find_project_root_prefers_git(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            (root / ".git").mkdir()
            nested = root / "a" / "b"
            nested.mkdir(parents=True)
            self.assertEqual(find_project_root(nested), root.resolve())

    def test_init_workspace_stays_inside_project(self):
        with tempfile.TemporaryDirectory() as tmp:
            project = Path(tmp)
            result = init_workspace(project)
            workspace = Path(result["workspace"])
            self.assertEqual(workspace, project.resolve() / ".claude")
            self.assertTrue((workspace / "CLAUDE.md").exists())
            self.assertTrue((workspace / "memory" / "index.md").exists())

    def test_init_preserves_existing_file(self):
        with tempfile.TemporaryDirectory() as tmp:
            project = Path(tmp)
            target = project / ".claude" / "context.md"
            target.parent.mkdir(parents=True)
            target.write_text("custom", encoding="utf-8")
            init_workspace(project)
            self.assertEqual(target.read_text(encoding="utf-8"), "custom")


if __name__ == "__main__":
    unittest.main()
