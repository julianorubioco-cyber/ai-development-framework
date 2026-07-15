from pathlib import Path
import re
import unittest

ROOT = Path(__file__).resolve().parents[1]

class RepositoryTests(unittest.TestCase):
    def test_all_skills_have_valid_frontmatter(self):
        for skill_dir in (ROOT / "skills").iterdir():
            if not skill_dir.is_dir():
                continue
            text = (skill_dir / "SKILL.md").read_text(encoding="utf-8")
            match = re.match(r"^---\n(.*?)\n---\n", text, re.S)
            self.assertIsNotNone(match, skill_dir.name)
            fm = match.group(1)
            self.assertIn(f"name: {skill_dir.name}", fm)
            self.assertIn("description:", fm)
            self.assertIn("argument-hint:", fm)

    def test_project_memory_is_not_global(self):
        text = (ROOT / "skills/memory/SKILL.md").read_text(encoding="utf-8")
        self.assertIn("<raiz-do-projeto>/.claude/", text)
        self.assertIn("Nunca grave conhecimento do", text)

    def test_orchestrator_has_both_gates(self):
        text = (ROOT / "skills/implement/SKILL.md").read_text(encoding="utf-8")
        self.assertIn("Não altere código antes de aprovação explícita", text)
        self.assertIn("REVIEW_STATUS: CHANGES_REQUIRED", text)

    def test_template_contains_domain_index(self):
        self.assertTrue((ROOT / "templates/project-workspace/.claude/memory/index.md").exists())


    def test_foundation_documents_exist(self):
        required = [
            "docs/001-vision.md",
            "docs/002-principles.md",
            "docs/003-architecture.md",
            "docs/004-token-economy.md",
            "docs/005-terminology.md",
            "docs/006-release-process.md",
        ]
        for rel in required:
            self.assertTrue((ROOT / rel).exists(), rel)

    def test_workspace_has_context_budget(self):
        self.assertTrue(
            (ROOT / "templates/project-workspace/.claude/context-budget.md").exists()
        )

    def test_token_policy_preserves_safety(self):
        text = (ROOT / "docs/004-token-economy.md").read_text(encoding="utf-8")
        self.assertIn("Economia nunca justifica uma aprovação sem evidência", text)

if __name__ == "__main__":
    unittest.main()
