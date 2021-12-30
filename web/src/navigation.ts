import { currentPath, setCurrentPath } from "./index";
import { refresh } from "./refresh";

export function previousDirectory() {
    if (currentPath !== "/") {
        setCurrentPath(currentPath.split("/").slice(0, -1).join("/"));
        refresh();
    }
}

export function navigateDirectory(event: MouseEvent, dirName: string) {
    let newPath = currentPath
    if (currentPath.charAt(-1) !== "/") {
        newPath += "/";
    }
    newPath += dirName;
    setCurrentPath(newPath)
    refresh();
}