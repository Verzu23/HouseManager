import 'dart:html';
import 'package:angular/angular.dart';
import 'package:angular_app/src/services/rest_service.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_icon/material_icon.dart';
import 'package:angular_components/utils/browser/window/module.dart';
import 'package:angular_forms/angular_forms.dart';
import 'package:angular_router/angular_router.dart';

import '../route_paths.dart';
import '../routes.dart';

@Component(
    selector: 'overview',
    styleUrls: ['overview_component.css'],
    templateUrl: 'overview_component.html',
    directives: [
      routerDirectives,
      MaterialButtonComponent,
      MaterialIconComponent,
      formDirectives,
      NgFor,
      NgIf,
    ],
    providers: [windowBindings, ClassProvider(RestService)],
    exports: [RoutePaths, Routes])
class OverviewComponent implements OnInit, OnDeactivate {
  String login() => RoutePaths.overview.toUrl();
  final Router _router;

  RestService restService;
  WebSocket ws;

  String OverviewComponentUrl() => RoutePaths.overview.toUrl();
  OverviewComponent(this._router, this.restService);

  @override
  Future<void> ngOnInit() async {
    connectWS();
  }

  Future<void> connectWS() async {
    ws = WebSocket('ws://127.0.0.1:3300/v1/ws');
    ws.onOpen.listen((event) {
      print('ConnectionOpen');
    });
    ws.onClose.listen((event) {
      print("ONCLOSE EVENT");
      //reconnect();
    });
    ws.onError.listen((event) {
      print("ONERROR EVENT");
    });
    ws.onMessage.listen((MessageEvent e) {
      ws.sendString('thanks man');
    });
  }

  @override
  void onDeactivate(RouterState current, RouterState next) {
    ws.close();
    print("ADIOS");
  }
}
