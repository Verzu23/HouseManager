import 'dart:convert';
import 'dart:html';
import 'package:angular_app/src/model/current_user.dart';
import 'package:angular_router/angular_router.dart';
import 'package:http/http.dart';
import 'package:angular/core.dart';

@Injectable()
class RestService {
  //-------------LOCAL------------//
  static final _headers = {'Content-Type': 'application/json;charset=utf-8'};
  static const _baseUrl = 'http://127.0.0.1:3300/v1';
  static const _baseManagementUrl = 'http://127.0.0.1:3300/usrManagement';
  static const _baseIP = '127.0.0.1:4332';

  static const _subGroup = '/v1';

  static const getProbesUrl = '/probes';
  static const getDetailedValueUrl = '/detailedValue';
  /*static const getLoginUrl = "/login";
  static const getLogoutUrl = "/logout";
  static const getWtReferenceUrl = "/GetWtReference";
  static const changeFlapConfigUrl = "/ChangeFlapConfig";
  static const changeFieldUrl = "/ChangeField";
  static const changeValueUrl = "/ChangeValue";
  static const calcBalanceUrl = "/CalcBalance";
  static const usrListUrl = "/usrList";
  static const postUserUrl = "/postUser";
  static const deleteUserUrl = "/deleteUser";
  static const updateUserUrl = "/updateUser";
  static const updateAccountInfoUrl = "/updateAccountInfo";*/

  final Client _http;
  final Router _router;
  WebSocket webSocket;

  CurrentUser currentUser = CurrentUser();

  RestService(this._http, this._router) {}

  /*Future<int> login() async {
    try {
      final response = await _http.post(_baseUrl + getLoginUrl, headers: _headers, body: jsonEncode(currentUser));
      if (response.statusCode == 200) {
        var body = jsonDecode(response.body);
        currentUser.fromJson(body);
        currentUser.password = '';
      }

      return response.statusCode;
    } catch (e) {
      _handleError(e, 'login');
      return null;
    }
  }

  Future<int> logout() async {
    try {
      final response = await _http.post(_baseUrl + getLogoutUrl, headers: _headers, body: jsonEncode(currentUser));
      return response.statusCode;
    } catch (e) {
      _handleError(e, 'logout');
      return null;
    }
  }

  Future<Map<String, dynamic>> getManagementData() async {
    try {
      var accessString = "Bearer " + currentUser.authToken;
      _headers['Authorization'] = accessString;
      final response = await _http.get(_baseManagementUrl + usrListUrl, headers: _headers);
      if (response.statusCode == 200) {
        var body = jsonDecode(response.body);
        var userList = <NewUser>[];
        var roleList = <Role>[];
        body['userList'].forEach((e) {
          userList.add(NewUser.fromJson(e));
        });

        body['roleList'].forEach((e) {
          roleList.add(Role.fromJson(e));
        });
        return {"userList": userList, "roleList": roleList};
      } else {
        await _router.navigate(RoutePaths.login.toUrl());
        return null;
      }
    } catch (e) {
      _handleError(e, 'getManagementData');
      return null;
    }
  }

  Future<void> postUser(NewUser user) async {
    try {
      var accessString = "Bearer " + currentUser.authToken;
      _headers['Authorization'] = accessString;
      final response = await _http.post(_baseManagementUrl + postUserUrl, headers: _headers, body: jsonEncode(user));
      if (response.statusCode == 200) {
        //TODO ERROR CHECK
      } else {}
    } catch (e) {
      _handleError(e, 'postUser');
      return null;
    }
  }

  Future<int> deleteUser(NewUser user) async {
    try {
      var accessString = "Bearer " + currentUser.authToken;
      _headers['Authorization'] = accessString;
      var packet = {'idUser': user.id.toString()};
      var uri = Uri.http(_baseIP, _subGroupManagement + deleteUserUrl, packet);
      final response = await _http.delete(uri, headers: _headers);
      if (response.statusCode == 200) {
        //TODO ERROR CHECK
      } else {}
      return response.statusCode;
    } catch (e) {
      _handleError(e, 'deleteUser');
      return null;
    }
  }

  Future<int> updateUser(NewUser user) async {
    try {
      var accessString = "Bearer " + currentUser.authToken;
      _headers['Authorization'] = accessString;
      final response = await _http.patch(_baseManagementUrl + updateUserUrl, headers: _headers, body: jsonEncode(user));
      if (response.statusCode == 200) {
        //TODO ERROR CHECK
      } else {}
    } catch (e) {
      _handleError(e, 'postUser');
      return null;
    }
  }

  Future<Map<int, String>> updateAccountInfo(NewUser user) async {
    try {
      var accessString = "Bearer " + currentUser.authToken;
      _headers['Authorization'] = accessString;
      final response = await _http.patch(_baseManagementUrl + updateAccountInfoUrl, headers: _headers, body: jsonEncode(user));
      Map<int, String> body = Map<int, String>();
      body[response.statusCode] = response.body;
      return body;
    } catch (e) {
      _handleError(e, 'postUser');
      return null;
    }
  }*/

  /*WebSocket OpenWebSocket() {
    if (webSocket == null) {
      webSocket = WebSocket("ws://" + _baseIP + _subGroup + "/ws");
    } else if (webSocket != null) {
      if (webSocket.readyState == WebSocket.CLOSED) {
        webSocket = WebSocket("ws://" + _baseIP + _subGroup + "/ws");
      }
    }
    return webSocket;
  }*/

  /*Future<List<Probe>> getProbes() async {
    try {
      //var accessString = "Bearer " + currentUser.authToken;
      //_headers['Authorization'] = accessString;
      final response = await _http.get(_baseUrl + getProbesUrl, headers: _headers);
      List<Probe> probeList = <Probe>[];
      if (response.statusCode == 200) {
        var body = jsonDecode(response.body);
        body.forEach((probeInfo) {
          probeList.add(Probe.fromJson(probeInfo));
        });

        return probeList;
      } else {
        //await _router.navigate(RoutePaths.login.toUrl());
        return null;
      }
    } catch (e) {
      _handleError(e, 'getProbes');
      return null;
    }
  }

  Future<List<dynamic>> getDetailedValue() async {
    try {
      //var accessString = "Bearer " + currentUser.authToken;
      //_headers['Authorization'] = accessString;
      final response = await _http.get(_baseUrl + getDetailedValueUrl, headers: _headers);
      if (response.statusCode == 200) {
        var body = jsonDecode(response.body);
        return body;
      } else {
        //await _router.navigate(RoutePaths.login.toUrl());
        return null;
      }
    } catch (e) {
      _handleError(e, 'getDetailedValue');
      return null;
    }
  }*/

  Exception _handleError(dynamic e, String function) {
    print(function + ' throw this error: ' + e.toString()); // for demo purposes only
    return Exception('Server error; cause: $e');
  }
}
