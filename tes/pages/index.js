// IMPORTS START
import BlogList from '@/components/blog/Blog';
import getAllPosts from '@/components/blog/blog_api';
// IMPORTS END
export default function Home({ allPosts, config }) {
  return (
    <div>
      {/* CONTENT START */}
      <BlogList allPosts={allPosts} config={config} />
      {/* CONTEND END */}
    </div>
  );
}

// STATIC PROPS START
export const getStaticProps = async () => {
  const allPosts = getAllPosts();
  return {
    props: { allPosts },
  };
};
// STATIC PROPS END
