var shoeData; // this is all the "CURRENT" shoe data from the sqlite db
var checkData; // this is all the "CURRENT" shoe data from the sqlite db
var modalID; // the current navette modal that's open
var shoeParams; // the current navette modal that's open

$(document).ready(function() {
    getDistParam()
        .then(getShoeDist())
        .then(getShoeCheck())
        .then(writeButtonDistance())
        .then(initModalData())
        .then(initChangeListeners())
        .then(initCheckListeners())

        // $("#date").addClass("form-control");
});

initCheckListeners = () => {
    // submit check
    $("#submit-shoe-check").click(async function() {
        const Timestamp = $("#date").val(); // string
        const Distance = parseInt($("#distance").val()); // distance needs to be int
        const Notes = $("#notes").val(); // string
        const Wear = parseFloat($("#measurement").val()) // float

        const checkFormData = {
            Shuttle: shoeData[modalID].Shuttle,
            Distance: Distance === NaN ? "nil" : Distance,
            Timestamp: Timestamp === "" ? "empty" : Timestamp,
            Notes: Notes,
            Wear: Wear,
            UUID: checkData[modalID] === undefined ? "" : checkData[modalID].UUID,
        };

        // console.log({ checkFormData });
        postData("http://localhost:8001/checkshoes", checkFormData).then((data) => {
            console.log(data);
        });
    });
}

initChangeListeners = () => {
    // submit change
    $("#submit-shoe-change").click(async function() {
        const New_Change_Date = $("#date").val(); // string
        const New_Change_Distance = parseInt($("#distance").val()); // distance needs to be int
        const New_Change_Notes = $("#notes").val(); // string
        const changeFormData = {
            Shuttle: shoeData[modalID].Shuttle,
            New_Change_Distance: New_Change_Distance === NaN ? "nil" : New_Change_Distance,
            New_Change_Date: New_Change_Date === "" ? "empty" : New_Change_Date,
            New_Change_Notes: New_Change_Notes,
            Previous_Change_UUID: shoeData[modalID].UUID,
        };

        console.log({
            changeFormData
        });
        postData("http://localhost:8001/changeshoes", changeFormData)
        .then((data) => {
            console.log(data);
        });
    });
}

postData = async (url = "", data = {}) => {
    const response = await fetch(url, {
            method: "POST", // or 'PUT'
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        })
        .then((response) => {
          return response.json();
        })
        .catch((error) => {
            console.error("Error:", JSON.stringify(error));
            return error;
        });
    return response;
}

getDistParam = async () => {
    fetch("http://localhost:8001/distparam")
        .then((response) => response.json())
        .then((data) => {
            shoeParams = data;
            console.log({
                shoeParams
            });
        });
}

getShoeCheck = async () => {
    fetch("http://localhost:8001/checkshoes")
        .then((response) => response.json())
        .then((data) => {
            checkData = data;
            console.log({
                checkData
            });
            // populate the button value with the distance from the db
            const navButtons = $("[id^=N]");
            const n = Object.entries(navButtons);
            n.forEach(([key, value]) => {
                if (checkData[value.id] !== undefined) {
                    if (checkData[value.id].Check_Shoes === true) {
                      $("#" + value.id).removeClass("btn-success");
                      $("#" + value.id).addClass("btn-danger");
                    }else{
                      $("#" + value.id).addClass("btn-success");
                      $("#" + value.id).removeClass("btn-danger");
                    }
                }
            });
        });
}

getShoeDist = async () => {
    fetch("http://localhost:8001/dist")
        .then((response) => response.json())
        .then((data) => {
            shoeData = data;
            console.log({
                shoeData
            });
            // populate the button value with the distance from the db
            const navButtons = $("[id^=N]");
            const n = Object.entries(navButtons);
            n.forEach(([key, value]) => {
                if (value.id !== undefined) {
                    let lastChange = shoeData[value.id].Shoe_Travel;
                    $("#" + value.id).text(lastChange);
                    if (lastChange >= shoeParams.Check) {
                        $("#" + value.id).removeClass("btn-success");
                        $("#" + value.id).addClass("btn-danger");
                    }
                }
            });
        });
}

initModalData = async () => {
    // navette buttons send data to modal form
    $("#distanceContainer").on("click", "button", function() {
        modalID = this.id;
        $("#notes").val("");
        $("#notes").val("");
        $("#notes").attr("placeholder", "Put any useful information here");
        $("#footer").scrollTop(0);
        $("#title").text(modalID + " Collector Shoes");
        $("#measurement").attr("placeholder", shoeParams.Min_Shoe.toFixed(1) + "-" + shoeParams.Max_Shoe.toFixed(1));
        $("#shoe-travel").text(shoeData[modalID].Shoe_Travel + " km");
        $("#days-installed").text(
            shoeData[modalID].Days_Installed === "" ?
            "0" :
            shoeData[modalID].Days_Installed
        );
        $("#current-distance").text(shoeData[modalID].Shoes_Last_Distance);
        $("#last-updated").text(shoeData[modalID].Last_Updated);
        $("#change-distance").text(shoeData[modalID].Shoes_Change_Distance);
        $("#change-date").text(shoeData[modalID].Shoes_Last_Changed);
        $("#change-notes").text(shoeData[modalID].Notes);

        if (checkData[modalID] !== undefined) {
            $("#check-distance").text(checkData[modalID].Last_Check_Distance);
            $("#check-timestamp").text(checkData[modalID].Last_Check_Timestamp);
            $("#check-wear").text(checkData[modalID].Last_Check_Wear.toFixed(1));
            $("#last-check-distance-1500km").text(checkData[modalID].Distance_1500km);
            $("#check-notes").text(checkData[modalID].Last_Check_Notes);
        }


    $("#check-date", "#check-distance", "#check-measurement", "#last-check-distance", "#last-check-distance-1500km", "#last-check-timestamp", "#last-check-wear", "#last-check-notes").val("");
    });
}

writeButtonDistance = () => {
    const navDiv = $("[id$=div]");
    const d = Object.entries(navDiv);
    d.forEach(([key, value]) => {
        if (value.id !== undefined) {
            let txt = value.id;
            let name = txt.slice(0, 4);
            $("#" + value.id).text(name);
        }
    });
//     });
}

// Clear the values when the modal closes
$("#changeModal,#checkModal").on("hidden.bs.modal", function() {
    $(this)
        .find("input,select")
        .val("")
        .end()
        .find("input[type=checkbox], input[type=radio]")
        .prop("checked", "")
        .end();
    $("#check-date", "#check-distance", "#check-measurement").val("");
    $("#change-notes").val(
        "Performed By: \nShoe Distance(km): \nShoe Measurement: \nOther Notes: "
    );
    $("#check-date", "#check-distance", "#check-measurement", "#last-check-distance", "#last-check-distance-1500km", "#last-check-timestamp", "#last-check-wear", "#last-check-notes").val("");
    $("#check-notes").val(
        "Performed By: \nShoe Distance(km): \nOther Notes: "
    );

});
