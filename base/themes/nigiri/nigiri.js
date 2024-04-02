const Navbar = () => {
  return (
    <div>
      <h1>Themed Navbar</h1>
    </div>
  );
};

// Export theme object
const theme = {
  dark: true,
  light: true,
  themedComponents: {
    Navbar: Navbar,
  },
};

export default theme;
