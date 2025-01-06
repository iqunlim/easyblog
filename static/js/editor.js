"use strict";
function initialize(tags, postID) {
    console.log("Loading editor...");
    // Required because of go templating
    const tagsArray = tags.split("%").filter((tag) => tag !== "");
    const editor = new Editor(tagsArray, postID);
    initalizeFileButtons();
    // Setting up form actions for the editor class
    const form = document.getElementById("blog-information");
    if (form === null) {
        throw new Error("The Form Was never found");
    }
    form.addEventListener("submit", (e) => {
        e.preventDefault();
        editor.UpdatePost();
    });
}
function initalizeFileButtons() {
    const actualBtn = document.getElementById("blog-img");
    const fileChosen = document.getElementById("blog-img-file-chosen");
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
    tags;
    postID;
    tagsDiv;
    form;
    constructor(tags, postID) {
        this.tags = tags;
        this.postID = Number(postID);
        this.tagsDiv = new TagsDiv(this.tags);
        this.form = document.getElementById("blog-information");
    }
    async UpdatePost() {
        Promise.all([this.BlogPostDataPost()])
            .then(() => (window.location.href = "/admin"))
            .catch((error) => console.log(error));
    }
    async BlogPostDataPost() {
        const formData = new FormData(this.form);
        const dataObj = {
            tags: this.tagsDiv.tags,
            title: formData.get("title"),
            content: formData.get("content"),
            summary: formData.get("summary"),
        };
        const dataToSend = JSON.stringify(dataObj);
        const httpMethod = this.postID === 0 ? "POST" : "PUT";
        const response = await fetch("/admin/posts" + (this.postID > 0 ? "/" + this.postID : ""), {
            credentials: "same-origin",
            mode: "same-origin",
            method: httpMethod,
            headers: { "Content-Type": "application/json" },
            body: dataToSend,
        });
        if (response.ok && (response.status === 200 || response.status === 201)) {
            return Promise.resolve(response.json());
        }
        else {
            return Promise.reject(new Error(`BlogPostDataPost request failed, response code ${response.status}`));
        }
    }
    async BlogPostImagePost() { }
}
class TagsDiv {
    tags;
    tagsDivElement;
    constructor(tags) {
        this.tags = tags;
        this.tagsDivElement = document.getElementById("tags");
        this.setInnerHTML();
    }
    addTag() {
        let tag = prompt("Enter tag");
        if (tag !== null &&
            tag !== undefined &&
            tag !== "" &&
            !this.tags.includes(tag)) {
            this.tags.push(tag);
            this.setInnerHTML();
        }
    }
    delTag(tag) {
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
    createTagButton(tag) {
        let newTagElement = document.createElement("button");
        newTagElement.addEventListener("click", (e) => e.preventDefault());
        newTagElement.textContent = `- ${tag}`;
        newTagElement.addEventListener("click", this.delTag.bind(this, tag));
        return newTagElement;
    }
    createAddButton() {
        let newTagElement = document.createElement("button");
        newTagElement.addEventListener("click", (e) => e.preventDefault());
        newTagElement.className = "tag-add";
        newTagElement.textContent = " + ";
        newTagElement.addEventListener("click", this.addTag.bind(this));
        return newTagElement;
    }
}
//# sourceMappingURL=editor.js.map