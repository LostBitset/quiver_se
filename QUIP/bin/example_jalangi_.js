J$.iids = {"8":[20,9,20,19],"9":[5,28,5,29],"10":[20,9,20,19],"16":[30,9,30,15],"17":[5,28,5,29],"18":[21,13,21,18],"25":[5,22,5,30],"26":[30,9,30,15],"33":[5,1,5,32],"41":[5,1,5,32],"49":[5,1,5,32],"57":[11,14,11,22],"65":[11,14,11,22],"73":[11,14,11,22],"81":[12,14,12,22],"89":[12,14,12,22],"97":[12,14,12,22],"105":[14,9,14,15],"113":[14,9,14,15],"121":[14,9,14,15],"129":[17,2,17,40],"137":[17,2,17,41],"145":[18,2,18,48],"153":[18,2,18,48],"161":[20,9,20,10],"169":[20,13,20,19],"177":[21,13,21,14],"185":[21,17,21,18],"193":[21,13,21,18],"201":[21,9,21,19],"209":[22,9,22,17],"217":[22,9,22,19],"225":[22,9,22,20],"233":[16,1,24,2],"241":[16,1,24,2],"249":[27,2,27,41],"257":[27,2,27,42],"265":[28,2,28,40],"273":[28,2,28,40],"281":[30,9,30,10],"289":[30,14,30,15],"297":[31,15,31,20],"305":[31,15,31,20],"313":[31,9,31,21],"321":[33,5,33,12],"329":[33,5,33,14],"337":[33,5,33,15],"345":[26,1,34,2],"353":[26,1,34,2],"361":[36,1,36,8],"369":[36,1,36,10],"377":[36,1,36,11],"385":[9,1,39,2],"393":[9,1,39,2],"401":[9,1,39,2],"409":[9,1,39,2],"417":[9,1,39,2],"425":[16,1,24,2],"433":[9,1,39,2],"441":[26,1,34,2],"449":[9,1,39,2],"457":[44,2,44,8],"465":[44,2,44,10],"473":[44,2,44,11],"481":[46,2,46,8],"489":[46,9,46,10],"497":[46,2,46,11],"505":[46,2,46,12],"513":[45,3,47,2],"521":[1,1,50,1],"529":[5,1,5,32],"537":[1,1,50,1],"545":[9,1,39,2],"553":[1,1,50,1],"561":[5,1,5,32],"569":[5,1,5,32],"577":[20,5,23,6],"585":[16,1,24,2],"593":[16,1,24,2],"601":[30,5,32,6],"609":[26,1,34,2],"617":[26,1,34,2],"625":[9,1,39,2],"633":[9,1,39,2],"641":[1,1,50,1],"649":[1,1,50,1],"nBranches":4,"originalCodeFileName":"/home/landon/Documents/repos/__active__/quiver_se/QUIP/bin/example.js","instrumentedCodeFileName":"/home/landon/Documents/repos/__active__/quiver_se/QUIP/bin/example_jalangi_.js","code":"\n// This code has been instrumented by function_info/instrument.js\n\n// bgn decl-prefix (static)\nfunction _Q$xnH(e) { throw e; }\n// end decl-prefix (static)\n\n// bgn entry-point (has-script)\nfunction _Q$ent() {\n\nvar sym__x = \"X:Real\";\nvar sym__y = \"Y:Real\";\n\nvar z = sym__x;\n\nfunction onFirst() {\n\t\"!!MAGIC@js_concolic/src-range=64:153\";\n\t\"!!MAGIC@js_concolic/idents=z:sym__y:onSecond\"\n\t\n    if (z < sym__y) {\n        z = z + 1;\n        onSecond();\n    }\n}\n\nfunction onSecond() {\n\t\"!!MAGIC@js_concolic/src-range=155:238\";\n\t\"!!MAGIC@js_concolic/idents=z:onFirst\"\n\t\n    if (z == 2) {\n        throw 'oof';\n    }\n    onFirst();\n}\n\nonFirst();\n\n\n}\n// end entry-point (has-script)\n\n// bgn main-rescue (actual-entry-point)\ntry {\n\t_Q$ent();\n} catch (e) {\n\t_Q$xnH(e);\n}\n// end main-rescue (actual-entry-point)\n\n"};
jalangiLabel4:
    while (true) {
        try {
            J$.Se(521, '/home/landon/Documents/repos/__active__/quiver_se/QUIP/bin/example_jalangi_.js', '/home/landon/Documents/repos/__active__/quiver_se/QUIP/bin/example.js');
            function _Q$xnH(e) {
                jalangiLabel0:
                    while (true) {
                        try {
                            J$.Fe(33, arguments.callee, this, arguments);
                            arguments = J$.N(41, 'arguments', arguments, 4);
                            e = J$.N(49, 'e', e, 4);
                            throw J$.X1(25, J$.Th(17, J$.R(9, 'e', e, 0)));
                        } catch (J$e) {
                            J$.Ex(561, J$e);
                        } finally {
                            if (J$.Fr(569))
                                continue jalangiLabel0;
                            else
                                return J$.Ra();
                        }
                    }
            }
            function _Q$ent() {
                jalangiLabel3:
                    while (true) {
                        try {
                            J$.Fe(385, arguments.callee, this, arguments);
                            function onFirst() {
                                jalangiLabel1:
                                    while (true) {
                                        try {
                                            J$.Fe(233, arguments.callee, this, arguments);
                                            arguments = J$.N(241, 'arguments', arguments, 4);
                                            J$.X1(137, J$.T(129, "!!MAGIC@js_concolic/src-range=64:153", 21, false));
                                            J$.X1(153, J$.T(145, "!!MAGIC@js_concolic/idents=z:sym__y:onSecond", 21, false));
                                            if (J$.X1(577, J$.C(8, J$.B(10, '<', J$.R(161, 'z', z, 0), J$.R(169, 'sym__y', sym__y, 0), 0)))) {
                                                J$.X1(201, z = J$.W(193, 'z', J$.B(18, '+', J$.R(177, 'z', z, 0), J$.T(185, 1, 22, false), 0), z, 0));
                                                J$.X1(225, J$.F(217, J$.R(209, 'onSecond', onSecond, 0), 0)());
                                            }
                                        } catch (J$e) {
                                            J$.Ex(585, J$e);
                                        } finally {
                                            if (J$.Fr(593))
                                                continue jalangiLabel1;
                                            else
                                                return J$.Ra();
                                        }
                                    }
                            }
                            function onSecond() {
                                jalangiLabel2:
                                    while (true) {
                                        try {
                                            J$.Fe(345, arguments.callee, this, arguments);
                                            arguments = J$.N(353, 'arguments', arguments, 4);
                                            J$.X1(257, J$.T(249, "!!MAGIC@js_concolic/src-range=155:238", 21, false));
                                            J$.X1(273, J$.T(265, "!!MAGIC@js_concolic/idents=z:onFirst", 21, false));
                                            if (J$.X1(601, J$.C(16, J$.B(26, '==', J$.R(281, 'z', z, 0), J$.T(289, 2, 22, false), 0)))) {
                                                throw J$.X1(313, J$.Th(305, J$.T(297, 'oof', 21, false)));
                                            }
                                            J$.X1(337, J$.F(329, J$.R(321, 'onFirst', onFirst, 0), 0)());
                                        } catch (J$e) {
                                            J$.Ex(609, J$e);
                                        } finally {
                                            if (J$.Fr(617))
                                                continue jalangiLabel2;
                                            else
                                                return J$.Ra();
                                        }
                                    }
                            }
                            arguments = J$.N(393, 'arguments', arguments, 4);
                            J$.N(401, 'sym__x', sym__x, 0);
                            J$.N(409, 'sym__y', sym__y, 0);
                            J$.N(417, 'z', z, 0);
                            onFirst = J$.N(433, 'onFirst', J$.T(425, onFirst, 12, false, 233), 0);
                            onSecond = J$.N(449, 'onSecond', J$.T(441, onSecond, 12, false, 345), 0);
                            var sym__x = J$.X1(73, J$.W(65, 'sym__x', J$.T(57, "X:Real", 21, false), sym__x, 1));
                            var sym__y = J$.X1(97, J$.W(89, 'sym__y', J$.T(81, "Y:Real", 21, false), sym__y, 1));
                            var z = J$.X1(121, J$.W(113, 'z', J$.R(105, 'sym__x', sym__x, 0), z, 1));
                            J$.X1(377, J$.F(369, J$.R(361, 'onFirst', onFirst, 0), 0)());
                        } catch (J$e) {
                            J$.Ex(625, J$e);
                        } finally {
                            if (J$.Fr(633))
                                continue jalangiLabel3;
                            else
                                return J$.Ra();
                        }
                    }
            }
            _Q$xnH = J$.N(537, '_Q$xnH', J$.T(529, _Q$xnH, 12, false, 33), 0);
            _Q$ent = J$.N(553, '_Q$ent', J$.T(545, _Q$ent, 12, false, 385), 0);
            try {
                J$.X1(473, J$.F(465, J$.R(457, '_Q$ent', _Q$ent, 1), 0)());
            } catch (e) {
                e = J$.N(513, 'e', e, 1);
                J$.X1(505, J$.F(497, J$.R(481, '_Q$xnH', _Q$xnH, 1), 0)(J$.R(489, 'e', e, 0)));
            }
        } catch (J$e) {
            J$.Ex(641, J$e);
        } finally {
            if (J$.Sr(649)) {
                J$.L();
                continue jalangiLabel4;
            } else {
                J$.L();
                break jalangiLabel4;
            }
        }
    }
// JALANGI DO NOT INSTRUMENT
