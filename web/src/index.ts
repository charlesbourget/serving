async function fetchDir(path: string) {
  const url = "http://localhost:8100" + path;
  const response = await fetch(url);
  if (response.status == 200) {
    const body = await response.text();
    return await JSON.parse(body);
  }
}

function createDirectoriesList(directories: Array<object>) {
  const list = document.getElementById("directories");
  cleanList(list);
  if (!directories) {
    return;
  }

  directories.forEach((directory) => {
    const node = document.createElement("li");
    const element = document.createElement("a");

    element.innerHTML = directory["name"];
    element.addEventListener("click", (event: MouseEvent) =>
      navigateDirectory(event, directory["name"])
    );
    node.appendChild(element);
    list.appendChild(node);
  });
}

function createFilesList(files: Array<object>) {
  const list = document.getElementById("files");
  cleanList(list);
  if (!files) {
    return;
  }

  files.forEach((file) => {
    const node = document.createElement("li");
    const element = document.createElement("a");

    element.setAttribute("href", file["link"]);
    element.innerHTML = file["name"];
    node.appendChild(element);
    list.appendChild(node);
  });
}

function cleanList(list: HTMLElement) {
  while (list.firstChild) {
    list.removeChild(list.firstChild);
  }
}

async function update() {
  document.getElementById("title").textContent = currentPath;
  const response = await fetchDir(currentPath);

  createDirectoriesList(response.directories);
  createFilesList(response.files);
}

function goUp() {
  if (currentPath !== "/") {
    currentPath = "/" + currentPath.split("/").slice(0, -1).join("/");
    update();
  }
}

function navigateDirectory(event: MouseEvent, dirName: string) {
  if (currentPath.charAt(-1) !== "/") {
    currentPath += "/";
  }
  currentPath += dirName;
  update();
}

function main() {
  update();
  document.getElementById("goUp").addEventListener("click", () => goUp());
}

let currentPath = "/.git";
window.onload = () => main();
