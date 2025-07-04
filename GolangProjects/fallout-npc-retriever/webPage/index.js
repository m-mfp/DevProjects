const npcName = document.getElementById("npcName");
const submitBtn = document.querySelector("input[type=submit]");
const result = document.querySelector(".result");

// Event Listener
submitBtn.addEventListener("click", async (e) => {
    e.preventDefault()

    if (npcName.value == ""){
        result.classList.add("hidden")
    } else {
        result.classList.remove("hidden")

        const npc = await getData(npcName.value);

        if (npc) {
            displayNPC(npc)
        } else {
            console.error("No NPC data found.")
        }
    }
});

// Request for the Golang API
async function getData(npcName) {
    const url = `http://localhost:12300/fallout-npc-scrapper/${encodeURIComponent(npcName)}`;
    const response = await fetch(url, {
        method: 'GET',
        cache: 'no-cache',
        headers: {
            'Content-Type': 'application/json'
        },
        referrerPolicy: 'no-referrer'
    });

    if (!response.ok) {
        // Handle errors
        console.error('Network response was not ok:', response.statusText);
        return;
    }

    const data = await response.json();
    return data
}


function displayNPC(npc) {
    for (let child of result.children) {
        if (child.nodeName == "H2") {
            child.innerText = npc.name
        } else {
            for (let info of child.children) {
                if (info.nodeName == "UL") {
                    const listItems = info.children
                    listItems[0].lastChild.innerText = npc.location.length > 40 ? npc.location.substring(0, 40) + '...' : npc.location;
                    listItems[1].lastChild.innerText = `${npc.essential ? 'Yes' : 'No'}`;
                    listItems[2].lastChild.innerText = `${npc.companion ? 'Yes' : 'No'}`;
                    listItems[3].lastChild.innerText = `${npc.merchant ? 'Yes' : 'No'}`;
                    listItems[4].lastChild.innerText = `${npc.doctor ? 'Yes' : 'No'}`;
                } else if (info.nodeName == "IMG") {
                    console.log(info)
                    info.src = 'https://static.wikia.nocookie.net/fallout/images/5/58/FO4_PrestonGarvey_Infobox.png'
                }
            }
        }
        
    }
}
