import React from 'react';
import styles from './Footer.module.css';

function Footer() {
  return (
    <footer className={styles.footer}>
      <div className={styles.footer__content}>
        <p>&copy; {new Date().getFullYear()} Harugo SSG</p>
      </div>
    </footer>
  );
}

export default Footer;
