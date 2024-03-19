// components/Layout.js

import Head from 'next/head';
import styles from './Layout.module.css';
import Navbar from '@/defaultComponents/navbar/Navbar';
import Footer from '@/defaultComponents/footer/Footer';
import { findPageName } from '@/utils/general';
import { useRouter } from 'next/router';

const Layout = ({ children, site_config, theme }) => {
  // const router = useRouter();
  // const pageName = findPageName(router.pathname, config.pages);

  return (
    <>
      <Head>
        {/* <title>{pageName ? pageName : config.site_title}</title> */}
        {/* <meta name="description" content={config.site_description} /> */}
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className={styles.container}>
        <Navbar
          site_config={site_config}

          // themedComponent={theme.themedComponents.Navbar}
        />
        <main>{children}</main>
        <Footer
          site_config={site_config}
          // themedComponent={theme.themedComponents.Footer}
        />
      </div>
    </>
  );
};

export default Layout;
