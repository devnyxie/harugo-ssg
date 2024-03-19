// IMPORTS START
import { Html } from 'next/document';
import Head from 'next/head';
import config_yml from '@/config/config.yaml';
import './defaultStyles/defaultStyles.css';
import Layout from '@/layouts/Layout';
// IMPORTS END

function App({ Component, pageProps }) {
  const config = config_yml ? config_yml : {};
  return (
    <>
      <Layout config={config}>
        <Component config={config} {...pageProps} />
      </Layout>
    </>
  );
}

export default App;
