import { join } from 'path';
import matter from 'gray-matter';
import fs from 'fs';

const postsDirectory = join('_posts');

// Ensure the directory exists, if not, create it
try {
  fs.statSync(postsDirectory);
} catch (error) {
  if (error.code === 'ENOENT') {
    // Directory doesn't exist, create it
    fs.mkdirSync(postsDirectory, { recursive: true });
  } else {
    throw error;
  }
}

export function getPostFileNames() {
  return fs.readdirSync(postsDirectory);
}

export function getPostByFileName(slug, fields) {
  const realSlug = slug.replace(/\.md$/, '');
  const fullPath = join(postsDirectory, `${realSlug}.md`);
  const fileContents = fs.readFileSync(fullPath, 'utf8');
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

export default function getAllPosts(
  fields = ['title', 'date', 'slug', 'coverImage', 'excerpt']
) {
  const slugs = getPostFileNames();
  const posts = slugs
    .map((slug) => getPostByFileName(slug, fields))
    .sort((post1, post2) => (post1.date > post2.date ? -1 : 1));
  return posts;
}
