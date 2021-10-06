import 'package:angular/angular.dart';
import 'package:angular_app/lang/EN.i69n.dart';
import 'package:angular_app/lang/IT.i69n.dart';
import 'package:angular_app/src/route_paths.dart';
import 'package:angular_app/src/routes.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';

@Component(
  selector: 'my-app',
  styleUrls: ['app_component.css'],
  templateUrl: 'app_component.html',
  directives: [
    routerDirectives,
    MaterialPersistentDrawerDirective,
    MaterialTemporaryDrawerComponent,
    DeferredContentDirective,
    MaterialButtonComponent,
    MaterialListComponent,
    MaterialListItemComponent,
    NgFor,
  ],
  providers: [],
  exports: [RoutePaths, Routes],
)
class AppComponent implements OnInit {
  AppComponent(this._router);

  var selectedLanguage;
  final Router _router;
  bool isMenuVisible = false;

  @override
  void ngOnInit() {
    selectedLanguage = EN();
  }

  void listSelection(String term, dynamic arg) {
    // do some validations...
    switch (term) {
      case 'accountSetting':
        {
          isMenuVisible = false;
          print('accountSetting');
        }
        break;
      case 'support':
        {
          isMenuVisible = false;
          print('support');
        }
        break;
      case 'sensorOverview':
        {
          _router.navigate(RoutePaths.overview.toUrl());
          isMenuVisible = false;
          print('sensorOverview');
        }
        break;
      case 'alarmSetting':
        {
          _router.navigate(RoutePaths.alarms.toUrl());
          isMenuVisible = false;
          print('alarmSetting');
        }
        break;
      case 'logout':
        {
          isMenuVisible = false;
          print('logout');
        }
        break;
      default:
    }
  }

  void isMenuVisibleChange(bool visible) {
    isMenuVisible = visible;
  }
}
