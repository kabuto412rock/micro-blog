var csrfTokenContent1 = document.getElementsByName("X-CSRF-TOKEN")[0].getAttribute("content");
var allForms = document.getElementsByTagName("form")
csrfTokenInput = document.createElement("input");
csrfTokenInput.name = "_csrf";
csrfTokenInput.type="hidden";
csrfTokenInput.value = csrfTokenContent1;
for (var i = 0; i < allForms.length; i++) {
  allForms[i].appendChild(csrfTokenInput.cloneNode());
}