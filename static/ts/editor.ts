interface blogPostData {
  content: FormDataEntryValue | null;
  image?: File | null;
  summary: FormDataEntryValue | null;
  tags: string[] | null;
  title: FormDataEntryValue | null;
}

function initialize(tags: string, postID: string) {
  console.log("Loading editor...");

  // Required because of go templating
  const tagsArray = tags.split("%").filter((tag) => tag !== "");
  const editor = new Editor(tagsArray, postID);
  initalizeFileButtons();
  // Setting up form actions for the editor class
  const form = <HTMLFormElement>document.getElementById("blog-information");
  if (form === null) {
    throw new Error("The Form Was never found");
  }
  form.addEventListener("submit", (e) => {
    e.preventDefault();
    editor.UpdatePost();
  });
}

function initalizeFileButtons() {
  const actualBtn = <HTMLInputElement>document.getElementById("blog-img")!;

  const fileChosen = document.getElementById("blog-img-file-chosen")!;

  if (fileChosen === null) {
  }

  actualBtn.addEventListener("change", function () {
    if (this.files === null) {
      throw new Error("Error in Image Upload Button, this.files is null");
    }
    fileChosen.textContent = this.files[0].name;
  });
}

class Editor {
  tags: string[];
  postID: number;
  tagsDiv: TagsDiv;
  constructor(tags: string[], postID: string) {
    this.tags = tags;
    this.postID = Number(postID);
    this.tagsDiv = new TagsDiv(this.tags);
  }

  UpdatePost() {
    const form = <HTMLFormElement>document.getElementById("blog-information");
    const formData = new FormData(form);
    const dataObj: blogPostData = {
      tags: this.tagsDiv.tags,
      title: formData.get("title"),
      content: formData.get("content"),
      summary: formData.get("summary"),
    };
    //Image data will have to be sent in its own fetch and returned with a url
    // if image name is null then set to default image (where to store this config?)
    // else set to /fs/<image uuid>.<image type>
    // before? that seems quite complicated...
    const dataToSend = JSON.stringify(dataObj);
    // TODO: Verify non-nulls and throw errors
    const httpMethod = this.postID === 0 ? "POST" : "PUT";
    fetch("/admin/posts" + (this.postID > 0 ? "/" + this.postID : ""), {
      credentials: "same-origin",
      mode: "same-origin",
      method: httpMethod,
      headers: { "Content-Type": "application/json" },
      body: dataToSend,
    }).then(() => (window.location.href = "/admin"));
  }
}

class TagsDiv {
  tags: string[];
  tagsDivElement: HTMLElement;
  constructor(tags: string[]) {
    this.tags = tags;
    this.tagsDivElement = document.getElementById("tags")!;
    this.setInnerHTML();
  }

  addTag() {
    let tag = prompt("Enter tag");

    if (
      tag !== null &&
      tag !== undefined &&
      tag !== "" &&
      !this.tags.includes(tag)
    ) {
      this.tags.push(tag);
      this.setInnerHTML();
    }
  }

  delTag(tag: string) {
    this.tags = this.tags.filter((oldTag) => oldTag !== tag);
    this.setInnerHTML();
  }

  setInnerHTML() {
    this.tagsDivElement.replaceChildren();
    this.tags.forEach((tag) => {
      this.tagsDivElement.appendChild(this.createTagButton(tag));
    });
    this.tagsDivElement.appendChild(this.createAddButton());
  }

  private createTagButton(tag: string) {
    let newTagElement = document.createElement("button");
    newTagElement.addEventListener("click", (e) => e.preventDefault());
    newTagElement.textContent = `- ${tag}`;
    newTagElement.addEventListener("click", this.delTag.bind(this, tag));
    return newTagElement;
  }

  private createAddButton() {
    let newTagElement = document.createElement("button");
    newTagElement.addEventListener("click", (e) => e.preventDefault());
    newTagElement.className = "tag-add";
    newTagElement.textContent = " + ";
    newTagElement.addEventListener("click", this.addTag.bind(this));
    return newTagElement;
  }
}
