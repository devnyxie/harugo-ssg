// IMPORTS START

import Navbar from '@/components/navbar/Navbar_component.jsx';

import Blog from '@/components/blog/Blog_component.jsx';

import blog_api from '@/components/blog/blog_api.js';
// IMPORTS END

export default function Page(props) {
  return (
    <div>
      {/* CONTENT START */}

      <Navbar {...props} />

      <Blog {...props} />
      {/* CONTENT END */}
    </div>
  );
}

// STATIC PROPS START
export async function getStaticProps() {
  const data = await blog_api();
  return { props: { blog_api: data } };
}
// STATIC PROPS END
