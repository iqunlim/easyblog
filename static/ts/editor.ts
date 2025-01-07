interface blogPostData {
  content: FormDataEntryValue | null;
  image?: File | null;
  imageUrl?: string | null;
  summary: FormDataEntryValue | null;
  tags: string[] | null;
  title: FormDataEntryValue | null;
}

interface APIResponse {
  code: number;
  message: string;
  success: boolean;
  data?: any;
}

interface APIImageResponse extends APIResponse {
  data: { filepath: string };
}

const DEFAULT_IMAGE = "/static/favicon.ico";
// API Response interface

function initialize(tags: string, postID: string) {
  console.log("Loading editor...");

  // Required because of go templating
  const tagsArray = tags.split("%").filter((tag) => tag !== "");
  const editor = new Editor(tagsArray, postID);
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

class Editor {
  postID: number;
  tagsDiv: TagsDiv;
  form: HTMLFormElement;
  currentImageUrl: string; /* Has the image changed... */
  imageBtn: HTMLInputElement;
  imagePreview: HTMLImageElement;
  constructor(tags: string[], postID: string) {
    this.postID = Number(postID);
    this.tagsDiv = new TagsDiv(tags);
    this.form = <HTMLFormElement>document.getElementById("blog-information");
    this.imageBtn = <HTMLInputElement>document.getElementById("blog-img")!;
    this.imagePreview = <HTMLImageElement>(
      document.getElementById("image-preview")
    );
    if (
      this.imagePreview.src === null ||
      this.imagePreview.src === undefined ||
      this.imagePreview.src === ""
    ) {
      this.currentImageUrl = DEFAULT_IMAGE;
    } else {
      this.currentImageUrl = this.imagePreview.src;
    }
    // Set up event listener for image previews
    this.form.addEventListener("change", this.setupImageElement.bind(this));

    //title input type = text name=title
    //content textarea id=blogbody name=content
    // div class= fileupload
    // input type = file name=image
    //imagePreview
    // summary textarea name=summary id=summary
    // tags already handled
  }

  async UpdatePost() {
    const formData = new FormData(this.form);

    // If we didnt update the image don't put to the image updater api
    if (
      this.currentImageUrl === this.imagePreview.src ||
      this.currentImageUrl === DEFAULT_IMAGE
    ) {
      this.BlogPostDataPost(formData, this.currentImageUrl)
        .then(() => (window.location.href = "/admin"))
        .catch((error) => console.log(error));
      // If we did then do so
    } else {
      this.BlogPostImagePost(formData)
        .then((response) =>
          this.BlogPostDataPost(formData, response.data.filepath)
        )
        .then(() => (window.location.href = "/admin"))
        .catch((error) => console.log(error));
    }
  }

  async BlogPostDataPost(
    formData: FormData,
    imageUrl: string | null
  ): Promise<APIResponse> {
    const dataObj: blogPostData = {
      tags: this.tagsDiv.tags,
      title: formData.get("title"),
      content: formData.get("content"),
      summary: formData.get("summary"),
      imageUrl: imageUrl,
    };
    if (dataObj.imageUrl === null || dataObj.imageUrl === "") {
      dataObj.imageUrl = DEFAULT_IMAGE;
    }
    const dataToSend = JSON.stringify(dataObj);
    const httpMethod = this.postID === 0 ? "POST" : "PUT";
    const response = await fetch(
      "/admin/posts" + (this.postID > 0 ? "/" + this.postID : ""),
      {
        credentials: "same-origin",
        mode: "same-origin",
        method: httpMethod,
        headers: { "Content-Type": "application/json" },
        body: dataToSend,
      }
    );
    if (response.ok && (response.status === 200 || response.status === 201)) {
      return response.json();
    } else {
      throw new Error(
        `BlogPostDataPost request failed, response code ${response.status}`
      );
    }
  }

  async BlogPostImagePost(
    formData: FormData
  ): Promise<APIImageResponse | APIResponse> {
    return await fetch("/admin/posts/image", {
      credentials: "same-origin",
      mode: "same-origin",
      method: "POST",
      body: formData,
    }).then((response) => {
      if (response.ok && response.status === 201) {
        return response.json();
      } else {
        throw new Error(
          `BlogPostImagePost request failed, response code ${response.status}`
        );
      }
    });
  }

  private setupImageElement() {
    if (
      this.imageBtn.files === null ||
      this.imageBtn.files[0] === null ||
      this.imageBtn.files[0] === undefined
    ) {
      return;
    }
    const file = this.imageBtn.files[0];
    const reader = new FileReader();
    reader.onload = (event) => {
      console.log(event.target);
      if (event.target === null || event.target.result === null) {
        throw new Error(
          "Error in setupImageElement, target or target result were null"
        );
      }
      this.imagePreview.src = event.target?.result.toString();
    };
    console.log(file);
    reader.readAsDataURL(file);
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
