library main;
import 'package:angular/angular.dart';
import 'package:di/di.dart';
import 'package:logging/logging.dart';
import 'package:angular/routing/module.dart';

// Temporary, please follow https://github.com/angular/angular.dart/issues/476
@MirrorsUsed(override: '*')
import 'dart:mirrors';

//読み込みモジュール定義
class MyModule extends Module {
  MyModule() {
    type(RankListCtrl);
    type(RankingService);
    factory(NgRoutingUsePushState,
        (_) => new NgRoutingUsePushState.value(false));
  }
}

//サービス
class RankingService {
  Http _http;
  
  RankingService(this._http);
  
  query() {
    return this._http.get("/objects").then((resp) {
      return resp.data;
    });
  }
}

@NgController(
  selector: '[rank-list]',
  publishAs: 'ctrl',
  map: const {
    'rankings': '=>rankings'
  }
)
class RankListCtrl {
  String title = "Hello Go world!!!";
  List<Map> rankings = [];
  
  RankingService _rankingService;

  RankListCtrl(this._rankingService){
    this._rankingService.query().then(
        (res) {
          this.rankings = res;
          print(this.rankings[0]);
        }
    );
  }
}

main(){
  Module myModule = new MyModule();
  ngBootstrap(module: myModule); 
}