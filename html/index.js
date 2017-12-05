window.Arduino = {};
window.onload = function() {
  Arduino.axios = axios.create({
    baseURL: 'http://localhost:8080/',
    timeout: 100000
  });

  // Arduino.axios = {
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
  //   put: function() {
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

  Arduino.chartsStatus = Arduino.initializeGoogleCharts();
  Arduino.initializeNavigation();
  Arduino.initializeDevices();
  Arduino.initDeviceDetail();
  Arduino.initNewDeviceDetail();
}

Arduino.initializeGoogleCharts = function() {
  google.charts.load('current', {'packages':['table']});

  return new Promise((resolve, reject) => {
    google.charts.setOnLoadCallback(() => {
      resolve();
    });
  });

}

Arduino.initializeNavigation = function() {

  $('.menu li').click((e) => {

    var selectedNav = $(e.currentTarget);
    $('.menu li').removeClass('active');
    selectedNav.addClass('active');

    var id = selectedNav.attr('data');
    $('.content-page').attr('hidden', 'hidden');
    $('.content-page[id="'+id+'"]').attr('hidden', false);
  });
}

Arduino.initializeDevices = function() {
  // Arduino.axios.setResponse({
  //   data: [
  //       {deviceId: 1, deviceName: 'Ovladac garaze',  deviceLocation: 'Garaz', deviceAddress: '192.168.0.44'},
  //       {deviceId: 2, deviceName: 'Bouda ovladac',   deviceLocation: 'Bouda',  deviceAddress: '192.168.0.41'},
  //       {deviceId: 3, deviceName: 'Pracovna',   deviceLocation: 'Dum 1.PP',  deviceAddress: '192.168.0.32'},
  //       {deviceId: 4, deviceName: 'Matejuv pokoj',   deviceLocation: 'Dum 1.NP',  deviceAddress: '192.168.0.46'}
  //     ]
  // });

  Arduino.axios.get('/devices/')
    .then(function (response) {
      Arduino.chartsStatus.then(() => {
        Arduino.drawDevicesTable(response.data);
      })
    })
    .catch(function (error) {
      console.log(error);
    });

    $('#devices .add').click(() => {
      $('#new-device-detail').removeClass('hidden');
    });
}

Arduino.drawDevicesTable = function(devices) {
  var data = new google.visualization.DataTable();
  data.addColumn('number', 'Id');
  data.addColumn('string', 'Name');
  data.addColumn('string', 'Location');
  data.addColumn('string', 'IP address');

  devices.forEach((d) => {
    data.addRow([d.deviceId, d.deviceName, d.deviceLocation, d.deviceLocation]);
  })

  var table = new google.visualization.Table(document.querySelector('#devices .table'));

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

  $('#device-detail .delete').click(() => {
    $('#device-detail').addClass('hidden');
    Arduino.axios.delete('/devices/' + $('#device-detail input[name="deviceId"]').val())
      .then(function (response) {
        console.log('deleted!')
      })
      .catch(function (error) {
        console.log(error);
      });
  });

  $('#device-detail .save').click(() => {
    $('#device-detail form').submit();
    $('#device-detail').addClass('hidden');
  });

  $('#device-detail form').submit((event) => {
    Arduino.axios.put('/devices/' + $('#device-detail input[name="deviceId"]').val(), {
        deviceId:  $('#device-detail input[name="deviceId"]').val(),
        deviceName:  $('#device-detail input[name="deviceName"]').val(),
        deviceLocation:  $('#device-detail input[name="deviceLocation"]').val(),
        deviceIP:  $('#device-detail input[name="deviceIP"]').val(),
        deviceType:  $('#device-detail input[name="deviceType"]').val(),
        deviceBoard:  $('#device-detail input[name="deviceBoard"]').val(),
        deviceSwVersion:  $('#device-detail input[name="deviceSwVersion"]').val(),
        targetServer:  $('#device-detail input[name="targetServer"]').val(),
        httpPort:  $('#device-detail input[name="httpPort"]').val(),
        note:  $('#device-detail input[name="note"]').val()

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

Arduino.initNewDeviceDetail = function() {
  $('#new-device-detail .cancel').click(() => {
    $('#new-device-detail').addClass('hidden');
  });

  $('#new-device-detail .save').click(() => {
    $('#new-device-detail form').submit();
    $('#new-device-detail').addClass('hidden');
  });

  $('#new-device-detail form').submit((event) => {
    Arduino.axios.post('/devices/' + $('#new-device-detail input[name="deviceId"]').val(), {
        deviceId:  $('#new-device-detail input[name="deviceId"]').val(),
        deviceName:  $('#new-device-detail input[name="deviceName"]').val(),
        deviceLocation:  $('#new-device-detail input[name="deviceLocation"]').val(),
      deviceIP:  $('#new-device-detail input[name="deviceIP"]').val(),
        deviceType:  $('#new-device-detail input[name="deviceType"]').val(),
        deviceBoard:  $('#new-device-detail input[name="deviceBoard"]').val(),
        deviceSwVersion:  $('#new-device-detail input[name="deviceSwVersion"]').val(),
        targetServer:  $('#new-device-detail input[name="targetServer"]').val(),
        httpPort:  $('#new-device-detail input[name="httpPort"]').val(),
        note:  $('#new-device-detail input[name="note"]').val()

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

Arduino.showDeviceDetail = function(deviceId) {
  // Arduino.axios.setResponse({
  //   "data": {
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
  // });

  Arduino.axios.get('/devices/' + deviceId)
    .then(function (response) {
      var device = response.data;

      $('#device-detail').removeClass('hidden');
      $('#device-detail input[name="deviceId"]').val(device.deviceId);
      $('#device-detail input[name="deviceName"]').val(device.deviceName);
      $('#device-detail input[name="deviceLocation"]').val(device.deviceLocation);
        $('#device-detail input[name="deviceIP"]').val(device.deviceIP);
      $('#device-detail input[name="deviceType"]').val(device.deviceType);
      $('#device-detail input[name="deviceBoard"]').val(device.deviceBoard);
      $('#device-detail input[name="deviceSwVersion"]').val(device.deviceSwVersion);
      $('#device-detail input[name="targetServer"]').val(device.targetServer);
      $('#device-detail input[name="httpPort"]').val(device.httpPort);
      $('#device-detail input[name="note"]').val(device.note);


    })
    .catch(function (error) {
      console.log(error);
    });
}
