import { currentPath } from "./index";
import { navigateDirectory } from "./navigation";

function createDirectoriesList(directories: Array<object>) {
  const list = document.getElementById("directories");
  cleanList(list);
  if (!directories) {
    return;
  }

  directories.forEach((directory) => {
    const node = document.createElement("div");
    const img = document.createElement("img");
    // @ts-ignore
    img.src = new URL(
      "../assets/directory.png?as=png&width=20&heigth=20",
      import.meta.url
    );
    node.appendChild(img);
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
    const node = document.createElement("div");
    const img = document.createElement("img");
    // @ts-ignore
    img.src = new URL(
      "../assets/file.png?as=png&width=20&heigth=20",
      import.meta.url
    );
    node.appendChild(img);

    const name = document.createElement("a");
    if (file["isText"]) {
      name.setAttribute(
        "href",
        `http://${window.location.hostname}:8100/static/viewer.html?${file["link"]}`
      );
    } else {
      name.setAttribute("href", file["link"]);
    }
    name.innerHTML = file["name"];
    node.appendChild(name);

    const size = document.createElement("span");
    size.innerHTML = file["size"] + " bytes";
    node.appendChild(size);

    list.appendChild(node);
  });
}

function cleanList(list: HTMLElement) {
  while (list.firstChild) {
    list.removeChild(list.firstChild);
  }
}

async function fetchDir(path: string) {
  const url = `http://${window.location.hostname}:8100/api${path}`;

  try {
    const response = await fetch(url);

    if (response.status == 200) {
      const body = await response.text();
      return await JSON.parse(body);
    } else {
      console.log(response.status);
      console.log("Error while fetching ressource");
    }
  } catch (error) {
    console.log(error);
  }
  return {};
}

export async function refresh() {
  document.getElementById("title").textContent = currentPath;
  document.getElementById("path").textContent = currentPath;
  const response = await fetchDir(currentPath);

  createDirectoriesList(response.directories);
  createFilesList(response.files);
}
