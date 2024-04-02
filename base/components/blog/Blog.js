import React, { useState } from 'react';
import { formatDateString } from '../../utils/general';
import Link from 'next/link';
import styles from './Blog.module.css';
import ComponentWrapper from '@/defaultComponents/componentWrapper/ComponentWrapper';

function Items({ currentItems, config }) {
  return (
    <>
      {currentItems.map((post, index) => (
        <div key={post.slug}>
          <div style={{ display: 'flex' }}>
            <div
              style={{
                width: 'max-content',
                whiteSpace: 'nowrap',
                marginRight: '0.5rem',
              }}
            >
              {formatDateString(post.date)}
            </div>
            <span
              style={{
                marginRight: '0.5rem',
              }}
            >
              {'>>'}
            </span>
            <div>
              <Link as={`/post/${post.slug}`} href="/post/[post]">
                <div className="fw-bold text-break">{post.title}</div>
              </Link>
            </div>
          </div>
        </div>
      ))}
    </>
  );
}

function PaginatedItems({ itemsPerPage, posts, theme, config }) {
  const isDark = theme === 'dark' ? true : false;
  const [page, setPage] = useState(1);
  const pageCount = Math.ceil(posts.length / itemsPerPage);
  const itemOffset = (page - 1) * itemsPerPage;
  const endOffset = itemOffset + itemsPerPage;
  const currentItems = posts.slice(itemOffset, endOffset);

  const handlePageChange = (event, value) => {
    setPage(value);
  };

  return (
    <>
      <Items currentItems={currentItems} config={config} />
      {/* {posts.length > config.max_posts_per_page ?  
      <Stack
        spacing={2}
        justifyContent="center"
        alignItems="center"
        mt={2}
        mb={2}
      >
        <Pagination
          shape="rounded"
          count={pageCount}
          page={page}
          onChange={handlePageChange}
          color="primary"
        />
      </Stack> : <></>} */}
    </>
  );
}

function BlogWrapper({ blog_api, theme, config }) {
  return (
    <div className={styles.Blog}>
      <h4 className="underlined_text">
        <div className="text">Blog</div>
      </h4>
      <div>
        <PaginatedItems
          config={config}
          itemsPerPage={
            config?.max_posts_per_page ? config?.max_posts_per_page : 8
          }
          posts={blog_api}
          theme={theme}
        />
      </div>
    </div>
  );
}

function Blog({ blog_api, theme, config }) {
  return (
    <ComponentWrapper theme={theme} config={config} funcName="blog">
      <BlogWrapper blog_api={blog_api} theme={theme} config={config} />
    </ComponentWrapper>
  );
}

export default Blog;
