const clipboard = document.getElementById("clipboard");

clipboard?.addEventListener("click", () => {
  const link = window.location.href;

  navigator.clipboard.writeText(link)

  const icon = clipboard.querySelector("i");
  icon?.classList.replace("fa-clipboard", "fa-check")
  clipboard.title = "copied"
})