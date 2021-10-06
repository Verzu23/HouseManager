import 'dart:async';
import 'dart:html';

import 'package:angular/angular.dart';

import 'package:angular_app/src/model/current_user.dart';
import 'package:angular_app/src/services/rest_service.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_router/angular_router.dart';
import 'package:angular_components/material_input/material_input.dart';
import 'package:angular_app/lang/EN.i69n.dart';
import '../route_paths.dart';
import '../routes.dart';

@Component(
    selector: 'login',
    styleUrls: ['login_component.css'],
    templateUrl: 'login_component.html',
    directives: [
      routerDirectives,
      MaterialInputComponent,
      MaterialButtonComponent,
      MaterialIconComponent,
      formDirectives,
      materialInputDirectives,
      NgFor,
      NgIf,
    ],
    providers: [ClassProvider(RestService), materialProviders],
    exports: [RoutePaths, Routes])
class LoginComponent implements OnInit {
  final Router _router;
  final RestService restService;
  String LogintUrl() => RoutePaths.login.toUrl();

  CurrentUser currentUser = CurrentUser();
  bool isChecked = false;
  var selectedLanguage;

  @ViewChild('#myInput1')
  HtmlElement usernameElement;

  @ViewChild('#myInput2')
  HtmlElement passwordElement;

  LoginComponent(this._router, this.restService);
  @override
  Future<Null> ngOnInit() async {
    selectedLanguage = EN();
    //selectedLanguage = IT();
    window.onStorage.listen((event) {
      window.localStorage.clear();
      _router.navigate(RoutePaths.login.toUrl());
    });

    currentUser = CurrentUser();
    currentUser.loadCurrentUser();
    if (currentUser.username != '' && currentUser.authToken != '') {
      onSelectLogin();
    }
  }

  Future<void> onSelectLogin() async {
    var message = '';
    showPopup(message);
    if (currentUser.username == '' || currentUser.password == '') {
      message = selectedLanguage.login.missUsrPas;
      showPopup(message);
    } else {
      var response;
      try {
        //response = await restService.login();
      } catch (e) {
        message = e.toString();
      }

      if (response == 200) {
        currentUser.saveCurrentUser();
        /* if (currentUser.rolename == "Administrator") {
          await _router.navigate(RoutePaths.userManager.toUrl());
        } else { */
        await _router.navigate(RoutePaths.login.toUrl());
        //}
      } else if (response == 401) {
        message = selectedLanguage.login.errUsrPas;
        currentUser.password = '';
        showPopup(message);
      } else {
        message = selectedLanguage.login.generic;
        showPopup(message);
      }
    }
  }

  void showPopup(String message) {
    var usernameElement = document.getElementById('user');
    var passwordElement = document.getElementById('pass');
    var popup = document.getElementById('myPopup');
    if (popup != null) {
      if (!popup.classes.contains('show') && message != '') {
        popup.classes.add('show');
      } else if (popup.classes.contains('show') && message == '') {
        popup.classes.remove('show');
      }
      popup.text = message;
      passwordElement.focus();
      usernameElement.focus();
    }
  }
}
