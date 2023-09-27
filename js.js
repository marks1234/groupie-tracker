//old js code, we are not actually using it anywhere

async function fetchData(endpoint) {
    try {
        console.log(endpoint)
        const response = await fetch(`http://localhost:8080${endpoint}`);
        console.log(response)
        const data = await response.text(); // Get the response body as text
        return JSON.parse(data); // Parse the text as JSON
    } catch (error) {
        console.error(`Error fetching ${endpoint}:`, error);
    }
}

async function loadArtists() {
    try {
        const artists = await fetchData('/artists');
        // const locations = await fetchData('/artists');
        const artistsContainer = document.getElementById('artists-container');
        console.log(artists)

        // Loop through the artists and display them
        artists.forEach(artist => {
            console.log(artist)
            // let responseObjects
            // document.addEventListener(window.onload, () => )
            const artistDiv = document.createElement('div');
            artistDiv.className = 'artist-card';

            artistDiv.innerHTML = `
            <div class="card">
                <div>
                    <img class="card-img" src="${artist.image}" alt="${artist.name}">
                    <h2 class="card-title">${artist.name}</h2>
                    <p class="card-info"><strong>Year Started:</strong> ${artist.creationDate}</p>
                    <p class="card-info"><strong>First Album:</strong> ${artist.firstAlbum}</p>
                    <p class="card-info"><strong>Members:</strong> ${artist.members}</p>
                </div>
            </div>
            `;
// replace artist.members and display a list of members instead
            artistsContainer.appendChild(artistDiv);
        });
    } catch (error) {
        console.error('Error loading data:', error);
    }
}

// Call the loadArtists function asynchronously
loadArtists();