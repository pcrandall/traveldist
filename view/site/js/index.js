var shoeData; // this is all the "CURRENT" shoe data from the sqlite db
var modalID; // the current navette modal that's open
var changeParams; // the current navette modal that's open

$(document).ready(function () {
  fetch("http://localhost:8001/distparam")
    .then((response) => response.json())
    .then((data) => {
      changeParams = data;
      console.log({ changeParams });
    });

  fetch("http://localhost:8001/dist")
    .then((response) => response.json())
    .then((data) => {
      shoeData = data;
      console.log({ shoeData });
      // populate the button value with the distance from the db
      const navButtons = $("[id^=N]");
      const n = Object.entries(navButtons);
      n.forEach(([key, value]) => {
        if (value.id !== undefined) {
          let dist = shoeData[value.id].Shoe_Travel;
          $("#" + value.id).text(dist);
          if (dist >= 1500) {
            $("#" + value.id).removeClass("btn-success");
            $("#" + value.id).addClass("btn-danger");
          }
        }
      });
    });

  // navette buttons send data to modal form
  $("#distanceContainer").on("click", "button", function () {
    modalID = this.id;
    $("#change-title").text(modalID + " Collector Shoe Change");
    $("#check-title").text(modalID + " Collector Shoes Check");
    $("#button-title").text(modalID + " Collector Shoes");
    $("#current-distance").text(shoeData[modalID].Shoes_Last_Distance);
    $("#last-updated").text(shoeData[modalID].Last_Updated);
    $("#last-distance").text(shoeData[modalID].Shoes_Change_Distance);
    $("#last-date").text(shoeData[modalID].Shoes_Last_Changed);
    $("#change-notes").val(
      "Performed By: \nShoe Distance(km): \nShoe Measurement: \nOther Notes: "
    );
    $("#check-notes").val(
      "Performed By: \nShoe Distance(km): \nOther Notes: "
    );
    $("#days-installed").text(
      shoeData[modalID].Days_Installed === ""
        ? "0"
        : shoeData[modalID].Days_Installed
    );
    $("#shoe-travel").text(shoeData[modalID].Shoe_Travel + " km");
    $("#change-last-notes").text(shoeData[modalID].Notes);
    $("#check-last-notes").text(shoeData[modalID].Notes);
  });

  const navDiv = $("[id$=div]");
  const d = Object.entries(navDiv);
  d.forEach(([key, value]) => {
    if (value.id !== undefined) {
      let txt = value.id;
      let name = txt.slice(0, 4);
      $("#" + value.id).text(name);
    }
  });

  // Clear the values when the modal closes
  $("#changeModal,#checkModal").on("hidden.bs.modal", function () {
    $(this)
      .find("input,select")
      .val("")
      .end()
      .find("input[type=checkbox], input[type=radio]")
      .prop("checked", "")
      .end();
    $("#change-notes").val(
      "Performed By: \nShoe Distance(km): \nShoe Measurement: \nOther Notes: "
    );
    $("#check-notes").val(
      "Performed By: \nShoe Distance(km): \nOther Notes: "
    );
  });

  // submit change
  $("#submit-shoe-change").click(async function () {
    const New_Change_Date = $("#change-date").val(); // string
    const New_Change_Distance = parseInt($("#change-distance").val()); // distance needs to be int
    const New_Change_Notes = $("#change-notes").val(); // string
    const changeFormData = {
      Shuttle: shoeData[modalID].Shuttle,
      New_Change_Distance:
        New_Change_Distance === NaN ? "nil" : New_Change_Distance,
      New_Change_Date: New_Change_Date === "" ? "empty" : New_Change_Date,
      New_Change_Notes: New_Change_Notes,
      Previous_Change_UUID: shoeData[modalID].UUID,
    };

    console.log({ changeFormData });
    postData("http://localhost:8001/changeshoes", changeFormData).then((data) => {
      console.log(data);
    });
  });
});

// submit check
$("#submit-shoe-check").click(async function () {
  const New_Check_Date = $("#check-date").val(); // string
  const New_Check_Distance = parseInt($("#check-distance").val()); // distance needs to be int
  const New_Check_Notes = $("#check-notes").val(); // string
  const New_Check_Measurement = parseFloat($("#check-measurement").val()) // string

  const checkFormData = {
    Shuttle: shoeData[modalID].Shuttle,
    New_Check_Distance:
      New_Check_Distance === NaN ? "nil" : New_Check_Distance,
    New_Check_Date: New_Check_Date === "" ? "empty" : New_Check_Date,
    New_Check_Notes: New_Check_Notes,
    New_Check_Measurement: New_Check_Measurement,
    Previous_Check_UUID: shoeData[modalID].UUID,
  };

  console.log({ checkFormData });
  postData("http://localhost:8001/checkshoes", checkFormData).then((data) => {
    console.log(data);
  });
});

async function postData(url = "", data = {}) {
  const response = await fetch(url, {
    method: "POST", // or 'PUT'
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log("Success:", data);
      return data;
    })
    .catch((error) => {
      console.error("Error:", JSON.stringify(error));
      return error;
    });
  return response;
}
