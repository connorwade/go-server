const form = {
  cont: document.querySelector("form"),
  title: document.querySelector("form #form-title"),
  artist: document.querySelector("form #form-artist"),
  price: document.querySelector("form #form-price"),
};

form.cont.addEventListener("submit", submitAlbum);

function submitAlbum(e) {
  const { title, artist, price } = form;
  e.preventDefault();
  const data = {
    title: title.nodeValue,
    artist: artist.nodeValue,
    price: price.nodeValue,
  };

  let jsonData = JSON.stringify(data);

  fetch("/newalbums", {
    method: "post",
    body: jsonData,
  });
}
