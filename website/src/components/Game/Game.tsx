import React from 'react';
// import logo from '../../assets/logo/logo512.png';
// import styles from './Game.module.css';
import SampleVis from '../visualisation/SampleVis';

function Home() {
  return (
    <div>
      {/* <img src={logo} className={styles.appLogo} alt="logo" /> */}
      <h2>Game Visualisation base</h2>
      <SampleVis data={[1, 2, 3]} />
    </div>
  );
}

export default Home;
