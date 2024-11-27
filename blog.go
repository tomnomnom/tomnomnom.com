package main

import (
	"embed"
	"errors"
	"html/template"
	"path"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed blog-posts
var blogPostContent embed.FS

const blogPostDir = "blog-posts"

type blogPost struct {
	ID      string
	Title   string
	Content template.HTML
}

type blogs []*blogPost

func (b blogs) find(id string) *blogPost {
	for _, p := range b {
		if p.ID == id {
			return p
		}
	}
	return nil
}

func getBlogPosts() (blogs, error) {
	out := make(blogs, 0)

	entries, err := blogPostContent.ReadDir(blogPostDir)
	if err != nil {
		return out, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		post, err := parseBlogPost(entry.Name())
		if err != nil {
			return out, err
		}

		out = append(out, post)
	}

	return out, nil
}

func parseBlogPost(filename string) (*blogPost, error) {
	raw, err := blogPostContent.ReadFile(path.Join(blogPostDir, filename))
	if err != nil {
		return nil, err
	}

	parts := strings.Split(filename, ".")
	if len(parts) != 2 {
		return nil, errors.New("blog post file lacks extension")
	}

	id := parts[0]
	format := parts[1]

	if format == "mkd" {
		raw = parseMarkdown(raw)
	}

	return &blogPost{
		ID: id,
		// TODO: parse the <h1> out of the page; also metadata?
		Title:   id,
		Content: template.HTML(raw),
	}, nil
}

func parseMarkdown(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
