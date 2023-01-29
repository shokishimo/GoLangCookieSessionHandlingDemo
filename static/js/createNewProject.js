document.addEventListener("DOMContentLoaded", function(){
    document.getElementById("add-another").addEventListener("click", function(){
        // Create the form fields
        const usernameLabel = document.createElement("label");
        usernameLabel.setAttribute("for", "username");
        usernameLabel.innerHTML = "Username:";
        const usernameInput = document.createElement("input");
        usernameInput.setAttribute("type", "text");
        usernameInput.setAttribute("id", "username");
        usernameInput.setAttribute("name", "username");

        const projectnameLabel = document.createElement("label");
        projectnameLabel.setAttribute("for", "projectname");
        projectnameLabel.innerHTML = "Project Name:";
        const projectnameInput = document.createElement("input");
        projectnameInput.setAttribute("type", "text");
        projectnameInput.setAttribute("id", "projectname");
        projectnameInput.setAttribute("name", "projectname");

        const osLabel = document.createElement("label");
        osLabel.setAttribute("for", "os");
        osLabel.innerHTML = "Operating System:";
        const osSelect = document.createElement("select");
        osSelect.setAttribute("id", "os");
        osSelect.setAttribute("name", "os");

        const osOption1 = document.createElement("option");
        osOption1.setAttribute("value", "Windows");
        osOption1.innerHTML = "Windows";
        const osOption2 = document.createElement("option");
        osOption2.setAttribute("value", "Mac");
        osOption2.innerHTML = "Mac";
        osSelect.appendChild(osOption1);
        osSelect.appendChild(osOption2);


        const browserLabel = document.createElement("label");
        browserLabel.setAttribute("for", "browser");
        browserLabel.innerHTML = "Browser:";
        const browserInput = document.createElement("input");
        browserInput.setAttribute("type", "text");
        browserInput.setAttribute("id", "browser");
        browserInput.setAttribute("name", "browser");

        const urlLabel = document.createElement("label");
        urlLabel.setAttribute("for", "url");
        urlLabel.innerHTML = "URL:";
        const urlInput = document.createElement("input");
        urlInput.setAttribute("type", "text");
        urlInput.setAttribute("id", "url");
        urlInput.setAttribute("name", "url");

        // Add line break elements
        const br1 = document.createElement("br");
        const br2 = document.createElement("br");
        const br3 = document.createElement("br");
        const br4 = document.createElement("br");
        const br5 = document.createElement("br");

        // Append the form fields to the form
        const form = document.getElementById("createNewForm");
        form.appendChild(usernameLabel);
        form.appendChild(usernameInput);
        form.appendChild(br1);
        form.appendChild(br2);
        form.appendChild(projectnameLabel);
        form.appendChild(projectnameInput);
        form.appendChild(br3);
        form.appendChild(br4);
        form.appendChild(osLabel);
        form.appendChild(osSelect);
        form.appendChild(br5);
        form.appendChild(br5);
        form.appendChild(browserLabel);
        form.appendChild(browserInput);
        form.appendChild(br5);
        form.appendChild(br5);
        form.appendChild(urlLabel);
        form.appendChild(urlInput);
        form.appendChild(br5);
        form.appendChild(br5);
    });

});