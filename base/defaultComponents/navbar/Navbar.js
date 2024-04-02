import Link from 'next/link';
import React from 'react';
import styles from './Navbar.module.css';

function Navbar(props) {
  const allPages = props.config.pages;
  // console.log(allPages);
  return (
    <div className={styles.Navbar}>
      <div className={styles.Navbar__content}>
        <div className={styles.Logo}>
          <Link
            href="/"
            style={{ alignSelf: 'self-start', textDecoration: 'none' }}
          >
            <h3 style={{ margin: 0, fontWeight: '400' }}>
              {props.config.site_title}
            </h3>
          </Link>
        </div>
        <div>
          {allPages &&
            allPages.length > 1 &&
            allPages.map((page) => {
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
