import React from 'react';

function ComponentWrapper({ children, theme, funcName }) {
  if (!theme || !children) {
    return;
  }
  if (theme.themedExists === false) {
    return <>{children}</>;
  }

  console.log(children);
  //   const funcName = children.type.name.toLowerCase();

  console.log('funcName: ', funcName);

  console.log('theme: ', theme);
  console.log(
    'theme.themedComponents[funcName]:',
    theme.themedComponents[funcName]
  );
  if (theme.themedComponents[funcName]) {
    console.log('themed loaded');
    const result = theme.themedComponents[funcName]();
    return result;
  } else {
    console.log('themed not loaded');

    return <>{children}</>;
  }
}

export default ComponentWrapper;
