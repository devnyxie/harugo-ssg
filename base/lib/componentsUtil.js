// componentsUtil.js
import path, { join } from 'path';
import fs from 'fs';

function capitalizeFirstChar(str) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

export function renderToStaticMarkup(element) {
  return new Promise((resolve, reject) => {
    let html = '';
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        html += chunk;
        callback();
      },
    });

    writableStream.on('finish', () => {
      resolve(html);
    });

    writableStream.on('error', reject);

    renderToPipeableStream(element).pipe(writableStream);
  });
}

async function getConfig() {
  const module = await import('@/config/site_config.yaml');
  return module.default;
}

async function getTheme() {
  const config = await getConfig();
  if (!config) {
    return null;
  }
  const module = await import(`@/theme/${config.theme}/${config.theme}.js`);
  return module.default;
}

export const findReactComponents = (html) => {
  const components = [];
  const re = /<([A-Z][a-zA-Z]+) /g;
  let match;
  while ((match = re.exec(html)) !== null) {
    components.push(match[1]);
  }
  return components;
};

const pagesDirectory = join('pages');

const extractThemedComponents = async ({ currentFileName }) => {
  let theme = await getTheme();
  if (!theme) {
    console.log('No theme found');
    return {};
  }
  const site_config = await getConfig();
  const fileName = path.basename(currentFileName);
  let currentPageName = fileName.split('.')[0];
  if (currentPageName === 'index') {
    for (let pageName in site_config.pages) {
      if (site_config.pages.hasOwnProperty(pageName)) {
        if (site_config.pages[pageName].index === 0) {
          currentPageName = pageName;
        }
      }
    }
  }

  // components from CLI
  const selectedComponentsViaCLI =
    site_config.pages[currentPageName].components;

  // console.log('selectedComponentsViaCLI:', selectedComponentsViaCLI);
  // components from Theme
  // scan the theme's "components" folder and get all the components (lowercased filename without extension).
  // theme's components path = `@/theme/${config.theme}/components`

  // const themeComponentsPath = path.resolve(
  //   `theme/${site_config.theme}/components`
  // );
  // // const themeComponentsPath = `../theme/`;
  // const themeComponents = fs.readdirSync(themeComponentsPath);
  // const themedComponents = themeComponents.map(
  //   (component) => component.split('.')[0]
  // );

  //
  const themedComponents = Object.keys(theme.themedComponents);

  // console.log('themedComponents:', themedComponents);
  // console.log(themedComponents);
  // const themedComponents = [];

  let themedComponentsObj = { themedExists: false };
  for (let [i, key] of Object.keys(selectedComponentsViaCLI).entries()) {
    // const component = selectedComponentsViaCLI[key];
    for (let j = 0; j < themedComponents.length; j++) {
      // if (themedComponents[j].toLowerCase() === component.name.toLowerCase()) {
      //   themedComponentsObj[component.name] = {
      //     name: component.name,
      //     index: component.index,
      //     path: `/theme/${site_config.theme}/components/${themedComponents[j]}.js`,
      //   };
      // }
      themedComponentsObj.themedExists = true;
    }

    // if no themed component was found for this component, then add it to the array.
    // first, check if such component.name already exists in the array

    // if (!themedComponentsObj[component.name]) {
    //   themedComponentsObj[component.name] = {
    //     name: component.name,
    //     index: component.index,
    //     path: `/components/${component.name}/${capitalizeFirstChar(
    //       component.name
    //     )}.js`,
    //   };
    // }
  }

  // console.log('final choice of components: ', themedComponentsObj);
  return themedComponentsObj;
};

export default extractThemedComponents;
