// componentsUtil.js
import App from 'next/app';
import { findAllComponentDefinitions } from 'react-dev-utils/findAllComponentDefinitions';

export const extractThemedComponents = (app) => {
  // Import the config dynamically
  const config = import('@/config/site_config.yaml').then((module) => {
    const config = module.default;
    return config;
  });
  // Import the theme dynamically
  theme = import(`@/theme/${config.theme}.js`).then((module) => {
    const theme = module.default;
    return theme;
  });
  //
  //
  const componentsSlice = {};
  const appSourceCode = app.toString();
  const components = findAllComponentDefinitions(appSourceCode);

  components.forEach((component) => {
    const componentName = component.displayName || component.name;
    if (theme.themedComponents.includes(componentName)) {
      componentsSlice[componentName] = theme.themedComponents[componentName];
    } else {
      componentsSlice[componentName] = componentName;
    }
  });

  return componentsSlice;
};
