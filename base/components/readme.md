# Harugo identifies components using suffixes and file extensions. The following are used:

## Suffixes

- `_api` for API components (Server sidee code, import and declaration will be created). Default exported function will be used as API handler. Currently only one API handler is allowed per file currently.
- `_slug` for slug components (Will be just moved to /pages/[slug].js). For example, `pages/post/[post].js` will be created for `post_slug.js`, which means that slug page will be accessible via `/post/some-post`.

## File Extensions

- `.css` for CSS components. We advise you to use modules. (Import will be created)

<strong>\*All other files will be considered as JS/JSX components (Import and declaration will be created).</strong>

# Themed Components

In order to add support of themed components for your component, please use <ThemedComponent funcName="funcName"> wrapper. `funcName` will be used to find themed component with identical name in the theme's config.
