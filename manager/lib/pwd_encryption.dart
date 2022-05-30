import 'dart:io';

import 'package:dbcrypt/dbcrypt.dart';

String encryptPassword(String password) {
  return DBCrypt().hashpw(password, DBCrypt().gensalt());
}

void saveEncryptedToFile(String password, File file) {
  String encryptedPassword = encryptPassword(password);
  file
    ..createSync(recursive: true)
    ..writeAsStringSync('{ "hash": "$encryptedPassword" }');
}
