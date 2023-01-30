document.addEventListener("DOMContentLoaded", function(){
    const appContainer = document.getElementById("addContainer");
    const addAppButton = document.getElementById("add-another");

    addAppButton.addEventListener("click", () => {
        // create a new app element
        const appElement = document.createElement("div");
        appElement.innerHTML += `
                <p class="app head">App</p>
                <label for="os">Operating System: </label>
                <select id="os" name="os" required>
                    <option value="Windows">Mac</option>
                    <option value="Mac">Windows</option>
                </select><br><br>
                <label for="browser">Browser Name (if it's a web app): </label>
                <input type="text" id="browser" name="browser" size="25" placeholder="otherwise, keep this blank" required><br><br>
                <label for="url">URL: </label>
                <input type="text" id="url" name="url" placeholder="or path to the app" required><br><br>
                <button class="delete-app-button">Delete</button><br><br><hr>
        `
        // add the new app element to the app container
        appContainer.appendChild(appElement);

        // add event listener for delete button
        const deleteButton = appElement.querySelector(".delete-app-button");
        deleteButton.addEventListener("click", () => {
            appElement.remove();
        });
    });
});