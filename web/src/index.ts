import { refresh } from './refresh'
import { previousDirectory } from './navigation';

export let currentPath = "/.git";

function main() {
  refresh();
  document.getElementById("goUp").addEventListener("click", () => previousDirectory());
}

export function setCurrentPath(newPath: string) {
  currentPath = newPath
}

window.onload = () => main();
