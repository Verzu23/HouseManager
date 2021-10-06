class Role {
  int id;
  int roleId;
  String rolename;

  Role({this.id, this.roleId, this.rolename});

  Role.fromJson(Map<String, dynamic> json) {
    id = json['Id'];
    roleId = json['RoleId'];
    rolename = json['Rolename'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    data['Id'] = this.id;
    data['RoleId'] = this.roleId;
    data['Rolename'] = this.rolename;
    return data;
  }
}
