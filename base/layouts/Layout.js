// components/Layout.js

import Head from 'next/head';
import Navbar from '@/defaultComponents/navbar/Navbar';
import Footer from '@/defaultComponents/footer/Footer';
import { findPageName } from '@/utils/general';
import { useRouter } from 'next/router';

const Layout = ({ children, site_config, theme }) => {
  const router = useRouter();
  // const pageName = findPageName(router.pathname, site_config.pages);
  return (
    <>
      <Head>
        {/* <title>{pageName ? pageName : site_config.site_title}</title> */}
        <meta name="description" content={site_config.site_description} />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="container">
        <Navbar site_config={site_config} />
        <main>{children}</main>
        <Footer site_config={site_config} />
      </div>
    </>
  );
};

export default Layout;
