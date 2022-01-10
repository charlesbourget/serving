import hljs from "highlight.js";

async function main() {
  const params = window.location.search.substring(1).split("&");
  if (params.length > 1) {
    alert("Too many params");
    return;
  }

  const codeField = document.getElementById("code");
  codeField.innerHTML = await fetchFile(params[0]);
  hljs.highlightAll();
}

async function fetchFile(url: string): Promise<string> {
  try {
    const response = await fetch(url);

    if (response.status == 200) {
      const body = await response.text();
      return body;
    } else {
      console.log(response.status);
      console.log("Error while fetching ressource");
    }
  } catch (error) {
    console.log(error);
  }
  return "";
}

window.onload = () => main();
