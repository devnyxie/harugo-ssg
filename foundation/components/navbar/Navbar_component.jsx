import Link from 'next/link';
import React from 'react';

function Navbar(props) {
  const allPages = props.config.pages;
  return (
    <div className="navbar">
      <div className="logo">
        <Link href="/">{props.config.site_title}</Link>
      </div>
      <div>
        {allPages &&
          allPages.map((page) => {
            return (
              <Link
                key={page.pageName}
                href={page.pagePath}
                className="navbar__menu__item"
              >
                {page.pageName}
              </Link>
            );
          })}
      </div>
    </div>
  );
}

export default Navbar;
