window.Arduino = {};
window.onload = function() { //provede se po nahrátí stránky
   Arduino.axios = axios.create({
     baseURL: '/',
       //baseURL: 'http://localhost:8080/',
     timeout: 100000
   });

  // Arduino.axios = {   //testovací data
  //   get: function() {
  //     return new Promise((resolve, reject) => {
  //       resolve(this.response);
  //     });
  //   },
  //
  //   post: function() {
  //     return new Promise((resolve, reject) => {
  //       resolve("OK");
  //     });
  //   },
  //
  //   delete: function() {
  //     return new Promise((resolve, reject) => {
  //       resolve("OK");
  //     });
  //   },
  //
  //   setResponse: function(response) {
  //     this.response = response;
  //   }
  // };

  Arduino.chartsStatus = Arduino.initializeGoogleCharts(); //grafy
  Arduino.initializeNavigation();
  Arduino.initializeDevices();
  Arduino.initDeviceDetail();
}

Arduino.initializeGoogleCharts = function() {   //na pozadí nahrává
  google.charts.load('current', {'packages':['table']});

  return new Promise((resolve, reject) => {
    google.charts.setOnLoadCallback(() => {   //co se stane, až se grafy nahrají
      resolve();      //tak se spustí toto
    });
  });

}

Arduino.initializeNavigation = function() {   //navigace

  $('.menu li').click((e) => {          //naco se klikne, to se ozna

    var selectedNav = $(e.currentTarget);
    $('.menu li').removeClass('active');
    selectedNav.addClass('active');

    var id = selectedNav.attr('data');
    $('.content-page').attr('hidden', 'hidden');  //nastaví, které stranky-položky menu- jsou hidden
    $('.content-page[id="'+id+'"]').attr('hidden', false);
  });
}

Arduino.initializeDevices = function() {
  // Arduino.axios.setResponse({
  //   data: {
  //    	devices: [
  //       {deviceId: 1, deviceName: 'Ovladac garaze',  deviceLocation: 'Garaz', deviceAddress: '192.168.0.44'},
  //       {deviceId: 2, deviceName: 'Bouda ovladac',   deviceLocation: 'Bouda',  deviceAddress: '192.168.0.41'},
  //       {deviceId: 3, deviceName: 'Pracovna',   deviceLocation: 'Dum 1.PP',  deviceAddress: '192.168.0.32'},
  //       {deviceId: 4, deviceName: 'Matejuv pokoj',   deviceLocation: 'Dum 1.NP',  deviceAddress: '192.168.0.46'}
  //     ]
  //   }
  // });

  Arduino.axios.get('/mydevices/')      //volání Rest api GET
    .then(function (response) {
      Arduino.chartsStatus.then(() => {   //vrátí promis - až je to nahrané(naloadované)
        console.log(response)
          Arduino.drawDevicesTable(response.data); //response.data.devices
      })
    })
    .catch(function (error) {
      console.log(error);
    });


}

Arduino.drawDevicesTable = function(devices) {
  var data = new google.visualization.DataTable();
  data.addColumn('number', 'Id');
  data.addColumn('string', 'Name');
  data.addColumn('string', 'Location');
  data.addColumn('string', 'Address');

  devices.forEach((d) => {
    data.addRow([d.deviceID, d.deviceName, d.deviceLocation, d.deviceIP]);
  })

  var table = new google.visualization.Table(document.getElementById('devices'));

  google.visualization.events.addListener(table, 'select', selectHandler);

  function selectHandler(e) {
    var selection = table.getSelection();

    if (selection.length == 1) {
      console.log('A table row was selected',  data.getFormattedValue(selection[0].row, 0))
      Arduino.showDeviceDetail(data.getFormattedValue(selection[0].row, 0));
    }
  }

  table.draw(data, {width: '100%'});
}

Arduino.initDeviceDetail = function() {
  $('#device-detail .cancel').click(() => {
    $('#device-detail').addClass('hidden');
  });

  $('#device-detail .delete').click(() => {     //DELETE
    $('#device-detail').addClass('hidden');
    Arduino.axios.delete('/devices/' + $('#deviceID').val())
      .then(function (response) {
        console.log('deleted!')
      })
      .catch(function (error) {
        console.log(error);
      });
  });

  $('#device-detail .save').click(() => {     //POST
    $('#device-detail form').submit();
    $('#device-detail').addClass('hidden');
  });

  $('#device-detail form').submit((event) => {
    console.log(event);
    Arduino.axios.post('/devices/' + deviceID, {       //POST  // get('/devices/' + deviceId,
      data: {
        deviceID:  $('#deviceID').val(),
        deviceName:  $('#deviceName').val(),
        deviceLocation:  $('#deviceLocation').val(),
        type:  $('#type').val(),
        deviceBoard:  $('#deviceBoard').val(),
        deviceSwVersion:  $('#deviceSwVersion').val(),
        targetServer:  $('#targetServer').val(),
        httpPort:  $('#httpPort').val(),
        note:  $('#note').val(),
        deviceIP:  $('#deviceIP').val()
      }
    })
      .then(function (response) {
        console.log('created!');
      })
      .catch(function (error) {
        console.log(error);
      });
    event.preventDefault();
  });
}

Arduino.showDeviceDetail = function(deviceID) {
  // Arduino.axios.setResponse({
  //   "data": {"device":
  //   	{
	//        "deviceId":2,
	//        "deviceName":"GarageController",
  //        "deviceLocation":"Garage",
	//        "deviceIP":"192.168.0.44",
  //
  //       "type": "arduino",
	//       "deviceBoard":"RobotDyn Wifi D1R2",
  //       "deviceSwVersion":"2017-11-14",
  //
	//       "targetServer":"192.168.0.18",
	//       "httpPort":9090,
	//       "note":"super device"
	//     }
  //   }
  // });

  Arduino.axios.get('/devices/' + deviceID)
    .then(function (response) {
      var device = response.data.device;

      $('#device-detail').removeClass('hidden');
      $('#deviceID').val(device.deviceID);
      $('#deviceName').val(device.deviceName);
      $('#deviceLocation').val(device.deviceLocation);
      $('#type').val(device.type);
      $('#deviceBoard').val(device.deviceBoard);
      $('#deviceSwVersion').val(device.deviceSwVersion);
      $('#targetServer').val(device.targetServer);
      $('#httpPort').val(device.httpPort);
      $('#note').val(device.note);
      $('#deviceIP').val(device.deviceIP);

    })
    .catch(function (error) {
      console.log(error);
    });
}
