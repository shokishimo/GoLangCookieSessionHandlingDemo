let count = 2;
document.addEventListener("DOMContentLoaded", function(){
    document.getElementById("add-another").addEventListener("click", function(){
        // Create the form fields
        const p = document.createElement("p")
        p.innerHTML = "App" + count++;

        const osLabel = document.createElement("label");
        osLabel.setAttribute("for", "os");
        osLabel.innerHTML = "Operating System: ";
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
        browserLabel.innerHTML = "Browser Name (if it's a web app): ";
        const browserInput = document.createElement("input");
        browserInput.setAttribute("type", "text");
        browserInput.setAttribute("id", "browser");
        browserInput.setAttribute("name", "browser");
        browserInput.setAttribute("size", "25");
        browserInput.setAttribute("placeholder", "otherwise, keep this blank")

        const urlLabel = document.createElement("label");
        urlLabel.setAttribute("for", "url");
        urlLabel.innerHTML = "URL: ";
        const urlInput = document.createElement("input");
        urlInput.setAttribute("type", "text");
        urlInput.setAttribute("id", "url");
        urlInput.setAttribute("name", "url");
        urlInput.setAttribute("placeholder", "or path to the app")

        // Append the form fields to the form
        const form = document.getElementById("addDiv");
        form.appendChild(p)
        form.appendChild(osLabel);
        form.appendChild(osSelect);
        form.appendChild(document.createElement("br"));
        form.appendChild(document.createElement("br"));
        form.appendChild(browserLabel);
        form.appendChild(browserInput);
        form.appendChild(document.createElement("br"));
        form.appendChild(document.createElement("br"));
        form.appendChild(urlLabel);
        form.appendChild(urlInput);
        form.appendChild(document.createElement("br"));
        form.appendChild(document.createElement("br"));
    });

});