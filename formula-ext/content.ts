const inject = (filepath: string, tag: string) => {
  const element = document.getElementsByTagName(tag)[0];
  const script = document.createElement("script");
  script.setAttribute("type", "text/javascript");
  script.setAttribute("src", filepath);
  script.setAttribute("id", "formula-ext");
  element.appendChild(script);
};

const start = () => {
  inject(chrome?.runtime?.getURL("web_accessible_resources.js"), "body");
};

start();
