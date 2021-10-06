import 'dart:convert';
import 'dart:html';


class CurrentUser {
  static final CurrentUser _currentUser = CurrentUser._internal();

  factory CurrentUser() {
    return _currentUser;
  }
  CurrentUser._internal();

  String username = '';
  String name = '';
  String surname = '';
  String rolename = '';
  String password = '';
  String email = '';

  String authToken;

  void saveCurrentUser() {
    var jsonPacket = json.encode({
      'Username': _currentUser.username,
      'Name': _currentUser.name,
      'Surname': _currentUser.surname,
      'Rolename': _currentUser.rolename,
      'AuthToken': _currentUser.authToken,
      'Email': _currentUser.email,
    });
    window.localStorage['CURRENTUSER'] = jsonPacket.toString();
  }

  void loadCurrentUser() {
    try {
      var jsonPacket = json.decode(window.localStorage['CURRENTUSER']);
      _currentUser.username = jsonPacket['Username'];
      _currentUser.name = jsonPacket['Name'];
      _currentUser.surname = jsonPacket['Surname'];
      _currentUser.rolename = jsonPacket['Rolename'];
      _currentUser.authToken = jsonPacket['AuthToken'];
      _currentUser.email = jsonPacket['Email'];
    } catch (e) {
      _currentUser.username = '';
      _currentUser.name = '';
      _currentUser.surname = '';
      _currentUser.rolename = '';
      _currentUser.authToken = '';
      _currentUser.email = '';
    }
  }

  Map<String, dynamic> toJson() {
    final data = <String, dynamic>{};
    data['Username'] = username;
    data['Password'] = password;
    data['AuthToken'] = authToken;
    return data;
  }

  void fromJson(Map<String, dynamic> json) {
    this.username = json['Username'];
    this.name = json['Name'];
    this.surname = json['Surname'];
    this.email = json['Email'];
    this.rolename = json['Rolename'];
    this.authToken = json['Access_Token'];
  }
}
