// IMPORTS START
import site_config_yaml from '@/config/site_config.yaml';
import user_config_yaml from '@/config/site_config.yaml';
import './defaultStyles/defaultStyles.css';
import Layout from '@/layouts/Layout';
import { useEffect, useState } from 'react';

// THEME IMPORT START

// THEME IMPORT END

// IMPORTS END
console.log(site_config_yaml);

function App({ Component, pageProps }) {
  const site_config = site_config_yaml ? site_config_yaml : {};
  const user_config = user_config_yaml ? user_config_yaml : {};
  // THEME START
  const [theme, setTheme] = useState(null);

  useEffect(() => {
    if (!site_config_yaml.theme) {
      return;
    }
    // Import the theme dynamically
    import(
      `@/theme/${site_config_yaml.theme}/${site_config_yaml.theme}.js`
    ).then((module) => {
      const theme = module.default;
      setTheme(theme);
    });
  }, []);
  // THEME END
  return (
    <>
      <Layout theme={theme} site_config={site_config} user_config={user_config}>
        <Component
          theme={theme}
          site_config={site_config}
          user_config={user_config}
          {...pageProps}
        />
      </Layout>
    </>
  );
}

export default App;
