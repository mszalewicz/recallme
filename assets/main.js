function toggleVisibility(id) {
    var object = document.getElementById(id);
    var objectClass = object.className;

    if(objectClass.includes("hidden")) {
        object.classList.remove("hidden");
        object.classList.add("visible");
    } else {
        object.classList.remove("visible");
        object.classList.add("hidden");
    }
}