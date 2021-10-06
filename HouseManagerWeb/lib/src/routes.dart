import 'package:angular_router/angular_router.dart';

import 'route_paths.dart';
import 'login/login_component.template.dart' as login_template;
import 'overview/overview_component.template.dart' as overview_component;
import 'alarms/alarms_component.template.dart' as alarms_component;

export 'route_paths.dart';

class Routes {
  static final login = RouteDefinition(
    routePath: RoutePaths.login,
    component: login_template.LoginComponentNgFactory,
  );
  static final overview = RouteDefinition(
    routePath: RoutePaths.overview,
    component: overview_component.OverviewComponentNgFactory,
  );
  static final alarms = RouteDefinition(
    routePath: RoutePaths.alarms,
    component: alarms_component.AlarmsComponentNgFactory,
  );

  static final all = <RouteDefinition>[
    //newProject,
    login,
    overview,
    alarms,
    RouteDefinition.redirect(
      path: '',
      redirectTo: RoutePaths.overview.toUrl(),
    ),
  ];
}
