import Link from 'next/link';
import React from 'react';
import styles from './Navbar.module.css';

function Navbar({ site_config }) {
  // const allPages = site_config.pages;
  // console.log(allPages);
  const allPages = [];
  for (var pageKey in site_config.pages) {
    if (site_config.pages.hasOwnProperty(pageKey)) {
      console.log(site_config.pages[pageKey].name);
      allPages.push({
        name: site_config.pages[pageKey].name,
        path: site_config.pages[pageKey].path,
      });
    }
  }

  return (
    <div className={styles.Navbar}>
      <div className={styles.Navbar__content}>
        <div className={styles.Logo}>
          <Link
            href="/"
            style={{ alignSelf: 'self-start', textDecoration: 'none' }}
          >
            <h3 style={{ margin: 0, fontWeight: '400' }}>
              {site_config.projectName}
            </h3>
          </Link>
        </div>
        <div>
          {allPages &&
            allPages.length > 1 &&
            allPages.map((page) => {
              console.log(page);
              return (
                <Link
                  key={page.name}
                  href={page.path}
                  className={styles.navbar__menu__item}
                >
                  {page.name}
                </Link>
              );
            })}
        </div>
      </div>
    </div>
  );
}

export default Navbar;
