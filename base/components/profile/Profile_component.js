import React from 'react';
import styles from './Profile.module.css';
import Image from 'next/image';

function Profile_component({ config }) {
  return (
    <div className={styles.profile}>
      <Image
        src="/pfp.jpg"
        alt="profile-image"
        className={styles.profile__image}
        width={180}
        height={180}
      />
      <div className={styles.profile__userDataBlock}>
        <div className={styles.profile__nameText}>{config.user.name}</div>
        <div className={styles.profile__positionText}>
          {config.user.position}
        </div>
        <div className={styles.profile__bioText}>{config.user.bio}</div>
      </div>
    </div>
  );
}

export default Profile_component;
