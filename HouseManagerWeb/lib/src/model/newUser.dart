import 'dart:math';

import 'package:angular_app/src/model/role.dart';
import 'package:intl/intl.dart';

const _daysInMonth = [0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];

class NewUser {
  int id;
  String username;
  String password;
  int idRole;
  int userId;
  String name;
  String surname;
  String email;
  int idTeam;
  DateTime validityFrom;
  DateTime validityTo;
  Role role;

  NewUser([
    this.id,
    this.username,
    this.password,
    this.idRole,
    this.userId,
    this.name,
    this.surname,
    this.email,
    this.idTeam,
    this.validityFrom,
    this.validityTo,
    this.role,
  ]);

  NewUser.fromJson(Map<String, dynamic> json) {
    id = json['Id'];
    username = json['Username'];
    password = json['Password'];
    idRole = json['IdRole'];
    userId = json['UserId'];
    name = json['Name'];
    surname = json['Surname'];
    email = json['Email'];
    idTeam = json['IdTeam'];
    validityFrom = DateTime.parse(json['ValidityFrom']);
    validityTo = DateTime.parse(json['ValidityTo']);
    role = json['Role'] != null ? Role.fromJson(json['Role']) : null;
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    data['Id'] = this.id;
    data['Username'] = this.username;
    data['Password'] = this.password;
    data['IdRole'] = this.idRole;
    data['UserId'] = this.userId;
    data['Name'] = this.name;
    data['Surname'] = this.surname;
    data['Email'] = this.email;
    data['IdTeam'] = this.idTeam;
    data['ValidityFrom'] = this.validityFrom.toIso8601String();
    data['ValidityTo'] = this.validityTo.toIso8601String();
    if (this.role != null) {
      data['Role'] = this.role.toJson();
    }
    return data;
  }

  String formattedStartDate(String lang) {
    String date;
    try {
      if (validityFrom != null) {
        date = DateFormat('d MMMM y', lang).format(validityFrom).toString().toLowerCase();
      }
    } catch (e) {
      return '-';
    }
    return date;
  }

  String formattedEndDate(String lang) {
    String date;
    try {
      if (validityTo != null) {
        date = DateFormat('d MMMM y', lang).format(validityTo).toString().toLowerCase();
      }
    } catch (e) {
      return '-';
    }
    return date;
  }

  bool isValidUser() {
    var valid = true;
    if (username == '' || username == null) {
      valid = false;
    }
    if (password == '' || password == null) {
      valid = false;
    }
    if (name == '' || name == null) {
      valid = false;
    }
    if (surname == '' || surname == null) {
      valid = false;
    }
    if (email == '' || email == null) {
      valid = false;
    }
    if (role == null) {
      valid = false;
    }
    return valid;
  }

  void calcValidityDates(int monthOfValidity) {
    validityFrom = DateTime.now();
    validityTo = addMonths(validityFrom, monthOfValidity);
  }

  bool isLeapYear(int value) => value % 400 == 0 || (value % 4 == 0 && value % 100 != 0);

  int daysInMonth(int year, int month) {
    var result = _daysInMonth[month];
    if (month == 2 && isLeapYear(year)) result++;
    return result;
  }

  DateTime addMonths(DateTime dt, int value) {
    var r = value % 12;
    var q = (value - r) ~/ 12;
    var newYear = dt.year + q;
    var newMonth = dt.month + r;
    if (newMonth > 12) {
      newYear++;
      newMonth -= 12;
    }
    var newDay = min(dt.day, daysInMonth(newYear, newMonth));
    if (dt.isUtc) {
      return DateTime.utc(newYear, newMonth, newDay, dt.hour, dt.minute, dt.second, dt.millisecond, dt.microsecond);
    } else {
      return DateTime(newYear, newMonth, newDay, dt.hour, dt.minute, dt.second, dt.millisecond, dt.microsecond);
    }
  }
}
