export function formatDateString(inputDateString, format, customSeparator) {
  let options;
  let date;
  if (format === 'full_date') {
    options = { year: 'numeric', month: 'long', day: 'numeric' };
    const inputDate = new Date(inputDateString);
    date = inputDate.toLocaleDateString('en-US', options);
  } else {
    options = {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    };
    const inputDate = new Date(inputDateString);
    date = inputDate
      .toLocaleDateString('en-US', options)
      .replace(
        /(\d+)\/(\d+)\/(\d+)/,
        `$3${customSeparator ? customSeparator : '-'}$1${
          customSeparator ? customSeparator : '-'
        }$2`
      );
  }

  return date;
}

export function findPageName(pathname, pages) {
  let page = pages.find((page) => page.path === pathname);
  return page ? `${page.name}` : '';
}
