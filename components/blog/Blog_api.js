import { join } from 'path';
import matter from 'gray-matter';
import { readdirSync, readFileSync, statSync, mkdirSync } from 'fs';

const postsDirectory = join(process.cwd(), '_posts');

// Ensure the directory exists, if not, create it
try {
  statSync(postsDirectory);
} catch (error) {
  if (error.code === 'ENOENT') {
    // Directory doesn't exist, create it
    mkdirSync(postsDirectory, { recursive: true });
  } else {
    throw error;
  }
}

export function getPostFileNames() {
  return readdirSync(postsDirectory);
}

export function getPostByFileName(slug, fields) {
  const realSlug = slug.replace(/\.md$/, '');
  const fullPath = join(postsDirectory, `${realSlug}.md`);
  const fileContents = readFileSync(fullPath, 'utf8');
  const { data, content } = matter(fileContents);

  const items = {};

  // Ensure only the minimal needed data is exposed
  fields.forEach((field) => {
    if (field === 'slug') {
      items[field] = realSlug;
    }
    if (field === 'content') {
      items[field] = content;
    }

    if (typeof data[field] !== 'undefined') {
      items[field] = data[field];
    }
  });

  return items;
}

export function getAllPosts(fields) {
  const slugs = getPostFileNames();
  const posts = slugs
    .map((slug) => getPostByFileName(slug, fields))
    .sort((post1, post2) => (post1.date > post2.date ? -1 : 1));
  return posts;
}
