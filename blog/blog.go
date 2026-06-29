package blog

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"about-me/model"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// postsDir is the directory, relative to the working directory, that holds the
// markdown source for every blog post. Drop a new `.md` file in here (with the
// frontmatter block described below) and it becomes a post automatically.
const postsDir = "content/blog"

// md renders markdown to HTML. GFM gives us tables, strikethrough, autolinks
// and fenced code blocks, which the posts rely on.
var md = goldmark.New(goldmark.WithExtensions(extension.GFM))

// LoadPosts reads every markdown file in postsDir and returns the parsed posts
// sorted newest-first by date.
func LoadPosts() ([]model.Blog, error) {
	entries, err := os.ReadDir(postsDir)
	if err != nil {
		return nil, err
	}

	var posts []model.Blog
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		post, err := parseFile(filepath.Join(postsDir, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", entry.Name(), err)
		}
		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})
	return posts, nil
}

// GetPost returns the single post matching slug, or (nil, nil) if none exists.
func GetPost(slug string) (*model.Blog, error) {
	posts, err := LoadPosts()
	if err != nil {
		return nil, err
	}
	for i := range posts {
		if posts[i].Slug == slug {
			return &posts[i], nil
		}
	}
	return nil, nil
}

// parseFile reads a markdown file, splits off its frontmatter and renders the
// body to HTML. The slug is derived from the filename.
func parseFile(path string) (model.Blog, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return model.Blog{}, err
	}

	meta, body := splitFrontmatter(string(raw))

	var buf bytes.Buffer
	if err := md.Convert([]byte(body), &buf); err != nil {
		return model.Blog{}, err
	}

	slug := strings.TrimSuffix(filepath.Base(path), ".md")
	return model.Blog{
		Slug:        slug,
		Title:       meta["title"],
		Date:        meta["date"],
		Author:      meta["author"],
		Description: meta["description"],
		Content:     template.HTML(buf.String()),
	}, nil
}

// splitFrontmatter separates a leading `---`-delimited YAML-ish block of
// `key: value` pairs from the markdown body. If no frontmatter is present the
// whole input is treated as the body.
func splitFrontmatter(content string) (map[string]string, string) {
	meta := map[string]string{}
	content = strings.TrimLeft(content, "\uFEFF \t\r\n")

	if !strings.HasPrefix(content, "---") {
		return meta, content
	}

	// Drop the opening delimiter line, then find the closing one.
	rest := strings.TrimPrefix(content, "---")
	rest = strings.TrimLeft(rest, "\r\n")
	end := strings.Index(rest, "\n---")
	if end == -1 {
		return meta, content
	}

	block := rest[:end]
	body := rest[end+len("\n---"):]
	body = strings.TrimLeft(body, "\r\n")

	for _, line := range strings.Split(block, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		key = strings.TrimSpace(strings.ToLower(key))
		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)
		meta[key] = value
	}
	return meta, body
}
