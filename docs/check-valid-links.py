import re
import sys
import requests
from pathlib import Path

def extract_links_from_file(file_path):
    with open(file_path, "r", encoding="utf-8") as file:
        content = file.read()
        links = re.findall(r"\]\(([^),]+)\)", content)
    return links

def test_links(links):
    broken_links = []
    for link in links:
        try:
            response = requests.head(link, allow_redirects=True, timeout=10)
            if response.status_code != 200:
                broken_links.append(link)
        except requests.RequestException as e:
            broken_links.append(link)
    return broken_links

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python check_links.py /path/to/markdown/files")
        sys.exit(1)

    root_dir = sys.argv[1]
    markdown_files = []
    for md_file in Path(root_dir).rglob("*.md"):
        if "node_modules" not in md_file.parts:
            markdown_files.append(md_file)
    for md_file in markdown_files:
        print(f"Checking links in {md_file}...")
        links = extract_links_from_file(md_file)
        broken_links = test_links(links)
        if broken_links:
            print("Broken links found:")
            for link in broken_links:
                print(f"- {link}")
        else:
            print("No broken links found.")
        print()